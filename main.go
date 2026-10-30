package main

import (
	"fmt"
	"log"

	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/cli/handlers/users"
	"github.com/zwoooz/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Config: %+v\n", cfg)

	s := cli.State{
		Cfg: &cfg,
	}

	cmds := cli.NewCommands()

	users.Register(cmds)
}
