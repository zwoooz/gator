// Package util contains the command handlers for utilities
package util

import "github.com/zwoooz/gator/internal/cli"

// Register adds the util commands to the registry
func Register(cmds *cli.Commands) {
	cmds.Register(helpSpec(cmds))
}
