package feeds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
	"github.com/zwoooz/gator/internal/rss"
)

func addFeed(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("%w: expected 2 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	if s.Cfg.CurrentUserName == "" {
		return fmt.Errorf("%w: no user currently logged in", cli.ErrUsage)
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

	user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("user does not exist")
		}
		return fmt.Errorf("error looking up current user: %w", err)
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

	fmt.Println(feed)

	return nil
}
