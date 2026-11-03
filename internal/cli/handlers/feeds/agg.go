package feeds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
)

func agg(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("%w: expected 1 argument, found %v", cli.ErrUsage, len(cmd.Args))
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid interval provided: %w", err)
	}
	if timeBetweenRequests < 10*time.Second {
		return fmt.Errorf("invalid interval provided '%v', minimum is 10s", timeBetweenRequests)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error getting feed: %v\n", err)
		}
	}
}

func scrapeFeeds(s *cli.State) error {
	nextFeed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No feeds are currently tracked, waiting until next interval...")
			return nil
		}
		return fmt.Errorf("error getting feed information: %w", err)
	}

	params := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        nextFeed.ID,
	}

	_, err = s.DB.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	rssFeed, err := s.RssClient.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed '%s': %w", nextFeed.Url, err)
	}

	fmt.Printf("\nFrom %s:\n", nextFeed.Name)
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* %s\n", item.Title)
	}
	return nil
}
