// Package users contains the command handlers for user management
package users

import (
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func login(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected 1 argument, found %v", len(cmd.Args))
	}
	username := cmd.Args[0]

	err := s.Cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Active user is now: %s\n", username)

	return nil
}
