package util

import (
	"fmt"

	"github.com/zwoooz/gator/internal/cli"
)

// helpSpec builds the help command
// It is a constructor instead of a plain handler because it needs the registry,
// which Handler's signature does not provide
func helpSpec(c *cli.Commands) cli.Spec {
	return cli.Spec{
		Name:        "help",
		Category:    "Util",
		Usage:       "help | help [command]",
		Description: "Shows usage information",
		Handler: func(s *cli.State, cmd cli.Command) error {
			if len(cmd.Args) > 0 {
				spec, ok := c.Lookup(cmd.Args[0])
				if ok {
					fmt.Println("Usage:")
					fmt.Printf("	%-24s %s\n", spec.Usage, spec.Description)
					return nil
				}
			}

			fmt.Println("Usage:")
			fmt.Println("	gator <COMMAND> [ARGUMENTS]")
			var last string
			for _, spec := range c.Specs() {
				if spec.Category != last {
					fmt.Printf("\n%s:\n", spec.Category)
					last = spec.Category
				}
				fmt.Printf("	%-24s %s\n", spec.Usage, spec.Description)
			}
			return nil
		},
	}
}
