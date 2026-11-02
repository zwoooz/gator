package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func login(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%w: expected 1 argument, found %v", cli.ErrUsage, len(cmd.Args))
	}
	name := cmd.Args[0]

	user, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user '%s' does not exist", name)
		}
		return fmt.Errorf("error getting user: %w", err)
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Active user is now: %s\n", user.Name)

	return nil
}
