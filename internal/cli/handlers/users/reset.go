package users

import (
	"context"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func reset(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("expected 0 arguments, found %v", len(cmd.Args))
		// TODO: print help/usage here OR even better do that in main
	}

	// WARN: add confirmation here, too tired right now
	err := s.DB.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Delted all rows in users table")

	return nil
}
