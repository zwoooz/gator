package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/cli/handlers/users"
	"github.com/zwoooz/gator/internal/config"
	"github.com/zwoooz/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(db)

	s := cli.State{
		Cfg: &cfg,
		DB:  dbQueries,
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
