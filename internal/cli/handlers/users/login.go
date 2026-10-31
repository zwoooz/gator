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

	_, err := s.DB.GetUser(context.Background(), name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user '%s' does not exist", name)
		}
		return err
	}

	err = s.Cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("Active user is now: %s\n", name)

	return nil
}
