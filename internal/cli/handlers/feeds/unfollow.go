package feeds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
)

func unfollow(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%w: expected 1 argument, found %v", cli.ErrUsage, len(cmd.Args))
	}
	feedURL := cmd.Args[0]

	feed, err := s.DB.GetFeedByUrl(context.Background(), feedURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed '%s' does not exist", feedURL)
		}
		return fmt.Errorf("error getting feed information: %w", err)
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	unfollowed, err := s.DB.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}
	if unfollowed == 0 {
		return fmt.Errorf("you aren't following this feed")
	}

	fmt.Println("Unfollowed feed:")
	fmt.Printf("* Name: %s\n", feed.Name)
	fmt.Printf("* URL: %s\n", feed.Url)

	return nil
}
