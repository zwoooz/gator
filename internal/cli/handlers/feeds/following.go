package feeds

import (
	"context"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
)

func following(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%w: expected 0 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	followed, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting followed feeds: %w", err)
	}
	if len(followed) == 0 {
		fmt.Println("You are not following any feeds")
		return nil
	}

	fmt.Println("Curently followed feeds:")
	for _, followedRow := range followed {
		fmt.Printf("* %-24s %s\n", followedRow.FeedName, followedRow.FeedUrl)
	}

	return nil
}
