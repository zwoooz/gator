// Package users contains the command handlers for user management
package users

import "github.com/zwoooz/gator/internal/cli"

var specs = []cli.Spec{
	{
		Name:        "login",
		Usage:       "login [username]",
		Description: "Set the current active user",
		Handler:     login,
	},
	{
		Name:        "register",
		Usage:       "register [username]",
		Description: "Create a new user and log in as them",
		Handler:     register,
	},
	{
		Name:        "reset",
		Usage:       "reset",
		Description: "Delete all users (asks for confirmation)",
		Handler:     reset,
	},
	{
		Name:        "users",
		Usage:       "users",
		Description: "Lists every user in the database, shows who is currently logged in",
		Handler:     users,
	},
}

// Register adds the user management commands to the registry
func Register(cmds *cli.Commands) {
	for _, spec := range specs {
		spec.Category = "Users"
		cmds.Register(spec)
	}
}
