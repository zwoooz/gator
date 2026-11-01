package feeds

import (
	"context"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func agg(s *cli.State, cmd cli.Command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("%w: expected 1 argument, found %v", cli.ErrUsage, len(cmd.Args))
	// }

	const URL = "https://www.wagslane.dev/index.xml"

	feed, err := s.RssClient.FetchFeed(context.Background(), URL)
	if err != nil {
		return err
	}
	fmt.Println(feed)

	return nil
}
