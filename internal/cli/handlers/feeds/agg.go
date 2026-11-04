package feeds

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/database"
	"github.com/zwoooz/gator/internal/rss"
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

	markFeedFetchedParams := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        nextFeed.ID,
	}

	_, err = s.DB.MarkFeedFetched(context.Background(), markFeedFetchedParams)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	rssFeed, err := s.RssClient.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	postCount := 0
	for _, item := range rssFeed.Channel.Item {
		trimmedTitle := strings.TrimSpace(item.Title)
		trimmedLink := strings.TrimSpace(item.Link)
		if trimmedTitle == "" || trimmedLink == "" {
			fmt.Fprintf(os.Stderr, "Skipping an item from '%s' with no title or no URL...\n", nextFeed.Name)
			continue
		}

		createPostParams := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       trimmedTitle,
			Description: rss.ToNullString(item.Description),
			Url:         trimmedLink,
			FeedID:      nextFeed.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: rss.ToNullTime(item.PubDate),
		}

		inserted, err := s.DB.CreatePost(context.Background(), createPostParams)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error saving post %s: %v\n", item.Title, err)
			continue
		}
		postCount += int(inserted)
	}
	if postCount > 0 {
		fmt.Printf("Saved %d posts from '%s'\n", postCount, nextFeed.Name)
	} else {
		fmt.Printf("No new posts from '%s'\n", nextFeed.Name)
	}

	return nil
}
