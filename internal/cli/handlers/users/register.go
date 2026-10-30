package users

import "github.com/zwoooz/gator/internal/cli"

type entry struct {
	name    string
	handler cli.Handler
}

var entries = []entry{
	{name: "login", handler: login},
}

func Register(cmds *cli.Commands) {
	for _, e := range entries {
		cmds.Register(e.name, e.handler)
	}
}
