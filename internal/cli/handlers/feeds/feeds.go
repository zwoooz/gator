package feeds

import (
	"context"
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

func feeds(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("%w: expected 0 arguments, found %v", cli.ErrUsage, len(cmd.Args))
	}

	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}
	if len(feeds) == 0 {
		fmt.Println("no feeds present")
		return nil
	}

	fmt.Println("Current tracked feeds:")
	var lastUser string
	for _, feedRow := range feeds {
		if lastUser != feedRow.CreatedBy {
			fmt.Printf("\n%s:\n", feedRow.CreatedBy)
		}
		fmt.Printf("* %-24s %s\n", feedRow.Name, feedRow.Url)
		lastUser = feedRow.CreatedBy
	}

	return nil
}
