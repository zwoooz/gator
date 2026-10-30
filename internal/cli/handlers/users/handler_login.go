package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("expected 1 argument, found %v", len(cmd.args))
	}
	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("Active user is now: %s\n", username)

	return nil
}
