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

func follow(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%w: expected 1 argument, found %v", cli.ErrUsage, len(cmd.Args))
	}

	feedURL := cmd.Args[0]

	feedRow, err := s.DB.GetFeedByUrl(context.Background(), feedURL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed '%s' does not exist", feedURL)
		}
		return fmt.Errorf("error getting feed information: %w", err)
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
