package feeds

import "github.com/zwoooz/gator/internal/cli"

var specs = []cli.Spec{
	{
		Name:        "agg",
		Usage:       "agg [URL]",
		Description: "",
		Handler:     agg,
	},
}

// Register adds the user management commands to the registry
func Register(cmds *cli.Commands) {
	for _, spec := range specs {
		spec.Category = "Feeds"
		cmds.Register(spec)
	}
}
