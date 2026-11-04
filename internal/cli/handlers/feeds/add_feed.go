package feeds

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
	"github.com/zwoooz/gator/internal/rss"
)

func addFeed(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("%w: expected 2 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	name, feedURL := cmd.Args[0], cmd.Args[1]
	u, err := url.Parse(feedURL)
	if err != nil || u.Scheme != "http" && u.Scheme != "https" || u.Host == "" {
		return fmt.Errorf("%w: invalid URL provided", cli.ErrUsage)
	}
	ok := rss.ValidateURL(feedURL)
	if !ok {
		return fmt.Errorf("invalid URL provided")
	}

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feed, err := s.DB.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		pqErr := pq.As(err, pqerror.UniqueViolation)
		if pqErr != nil {
			return fmt.Errorf("feed already added, follow it instead")
		}
		return fmt.Errorf("error creating feed: %w", err)
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = s.DB.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Added feed '%s' with URL '%s'\n", feed.Name, feed.Url)

	return nil
}
