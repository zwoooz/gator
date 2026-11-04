// Package feeds contains the command handlers for feed management
package feeds

import "github.com/zwoooz/gator/internal/cli"

var specs = []cli.Spec{
	{
		Name:        "agg",
		Usage:       "agg [interval]",
		Description: "Fetches the current least updated feed on the given interval (e.g. 30s, 5m or 1h), minimum 10s, keeps running until stopped with Ctrl+C",
		Handler:     agg,
	},
	{
		Name:        "addfeed",
		Usage:       "addfeed [name] [URL]",
		Description: "Adds feed to the current logged in user",
		Handler:     cli.MiddlewareLoggedIn(addFeed),
	},
	{
		Name:        "feeds",
		Usage:       "feeds",
		Description: "List all tracked feeds and who created them",
		Handler:     feeds,
	},
	{
		Name:        "follow",
		Usage:       "follow [URL]",
		Description: "Follows an already tracked feed as the logged in user",
		Handler:     cli.MiddlewareLoggedIn(follow),
	},
	{
		Name:        "unfollow",
		Usage:       "unfollow [URL]",
		Description: "Unfollows a followed feed as the logged in user",
		Handler:     cli.MiddlewareLoggedIn(unfollow),
	},
	{
		Name:        "following",
		Usage:       "following",
		Description: "List all feeds followed by the logged in user",
		Handler:     cli.MiddlewareLoggedIn(following),
	},
	{
		Name:        "browse",
		Usage:       "browse | browse [limit]",
		Description: "Shows saved posts, with optional limit on amount of posts",
		Handler:     cli.MiddlewareLoggedIn(browse),
	},
}

// Register adds the feed management commands to the registry
func Register(cmds *cli.Commands) {
	for _, spec := range specs {
		spec.Category = "Feeds"
		cmds.Register(spec)
	}
}
