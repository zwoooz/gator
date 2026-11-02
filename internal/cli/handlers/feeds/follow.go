package feeds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
)

func follow(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%w: expected 1 argument, found %v", cli.ErrUsage, len(cmd.Args))
	}

	feedURL := cmd.Args[0]

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

	feedRow, err := s.DB.GetFeedByUrl(context.Background(), feedURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed does not exist")
		}
		return fmt.Errorf("error getting feed from database: %w", err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feedRow.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	followRow, err := s.DB.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		pqErr := pq.As(err, pqerror.UniqueViolation)
		if pqErr != nil {
			return fmt.Errorf("already following feed")
		}
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("User '%s' started following feed:\n", followRow.UserName)
	fmt.Printf("Name: %-24s\n", followRow.FeedName)
	fmt.Printf("URL: %-24s\n", feedRow.Url)

	return nil
}
