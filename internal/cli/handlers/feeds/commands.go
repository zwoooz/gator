// Package feeds contains the command handlers for feed management
package feeds

import "github.com/zwoooz/gator/internal/cli"

var specs = []cli.Spec{
	{
		Name:        "agg",
		Usage:       "agg",
		Description: "",
		Handler:     agg,
	},
	{
		Name:        "addfeed",
		Usage:       "addfeed [name] [URL]",
		Description: "Adds feed to the current logged in user",
		Handler:     addFeed,
	},
	{
		Name:        "feeds",
		Usage:       "feeds",
		Description: "List all tracked feeds and who created them",
		Handler:     feeds,
	},
	{
		Name:        "follow",
		Usage:       "follow [feed url]",
		Description: "Follows an already tracked feed as the logged in user",
		Handler:     follow,
	},
	{
		Name:        "following",
		Usage:       "following",
		Description: "List all feeds beign followed by the logged in user",
		Handler:     following,
	},
}

// Register adds the user management commands to the registry
func Register(cmds *cli.Commands) {
	for _, spec := range specs {
		spec.Category = "Feeds"
		cmds.Register(spec)
	}
}
