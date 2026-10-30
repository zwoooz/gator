package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/cli/handlers/users"
	"github.com/zwoooz/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := cli.State{
		Cfg: &cfg,
	}

	cmds := cli.NewCommands()

	users.Register(cmds)

	cmd, err := cli.ParseInput(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmds.Run(&s, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
