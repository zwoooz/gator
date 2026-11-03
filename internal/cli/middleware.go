package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/zwoooz/gator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		name := s.Cfg.CurrentUserName
		if name == "" {
			return fmt.Errorf("%w: no user currently logged in", ErrUsage)
		}

		user, err := s.DB.GetUser(context.Background(), name)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("user '%s' does not exist", name)
			}
			return fmt.Errorf("error getting user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
