package feeds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func following(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%w: expected 0 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	name := s.Cfg.CurrentUserName
	if name == "" {
		return fmt.Errorf("%w: no user currently logged in", cli.ErrUsage)
	}

	user, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user '%s' does not exist", name)
		}
		return fmt.Errorf("error getting user: %w", err)
	}

	followed, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting followed feeds: %w", err)
	}
	if len(followed) == 0 {
		fmt.Println("You are not following any feeds")
	}

	fmt.Println("Curently followed feeds:")
	for _, followedRow := range followed {
		fmt.Printf("* %s\n", followedRow.FeedName)
	}

	return nil
}
