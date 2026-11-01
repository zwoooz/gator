package users

import (
	"context"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func users(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%w: expected 0 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}
	if len(users) == 0 {
		return fmt.Errorf("no registered users")
	}

	for _, user := range users {
		if user == s.Cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
			continue
		}
		fmt.Printf("* %s\n", user)
	}

	return nil
}
