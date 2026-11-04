package feeds

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
)

func browse(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("%w: expected 0 or 1 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	limit := 2
	if len(cmd.Args) == 1 {
		parsed, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error setting limit '%s', defaulting to 2\n", cmd.Args[0])
			limit = 2
		} else {
			limit = parsed
		}
		if limit < 1 {
			fmt.Fprintf(os.Stderr, "negative limit '%d' used, setting limit to '1'\n", limit)
			limit = 1
		}
		if limit > 100 {
			fmt.Fprintf(os.Stderr, "Too large limit '%d' used, setting limit to '100'\n", limit)
			limit = 100
		}
	}

	getPostsForUserParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), getPostsForUserParams)
	if err != nil {
		return fmt.Errorf("error getting saved posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts to show...")
	}
	for i, post := range posts {
		fmt.Printf("\nPost %d out of %d:\n", i+1, len(posts))
		fmt.Printf("%s\n", post.Title)

		meta := post.FeedName
		if post.PublishedAt.Valid {
			meta += " - " + post.PublishedAt.Time.Local().Format("Jan 2, 2006 at 15:04 MST")
		}
		fmt.Println(meta)

		if post.Description.Valid {
			desc := strings.Join(strings.Fields(post.Description.String), " ")
			runes := []rune(desc)
			if len(runes) > 200 {
				desc = string(runes[:200]) + "..."
			}
			if !strings.HasPrefix(desc, "<") {
				fmt.Println(desc)
			}
		}

		fmt.Println(post.Url)
	}

	return nil
}
