package feeds

import (
	"context"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func agg(s *cli.State, cmd cli.Command) error {
	const URL = "https://www.wagslane.dev/index.xml"

	feed, err := s.RssClient.FetchFeed(context.Background(), URL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	fmt.Println(feed)

	return nil
}
