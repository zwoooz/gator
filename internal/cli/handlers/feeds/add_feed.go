package feeds

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
)

func addFeed(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 2 {
		// WARN: bootdev does not accept exit code 2, fix replace below with comment after
		// return fmt.Errorf("%w: expected 2 arguments, found %v", cli.ErrUsage, len(cmd.Args))
		return fmt.Errorf("expected 2 arguments, found %v", len(cmd.Args))
	}

	if s.Cfg.CurrentUserName == "" {
		// WARN: bootdev does not accept exit code 2, fix replace below with comment after
		// return fmt.Errorf("%w: no user currently logged in", cli.ErrUsage)
		return fmt.Errorf("no user currently logged in")
	}

	name, feedURL := cmd.Args[0], cmd.Args[1]
	u, err := url.Parse(feedURL)
	if err != nil || u.Scheme != "http" && u.Scheme != "https" || u.Host == "" {
		// WARN: bootdev does not accept exit code 2, fix replace below with comment after
		// return fmt.Errorf("%w: invalid URL provided", cli.ErrUsage)
		return fmt.Errorf("invalid URL provided")
	}

	user, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error looking up current user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       feedURL,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feed, err := s.DB.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	fmt.Println(feed)

	return nil
}
