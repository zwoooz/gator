package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/zwoooz/gator/internal/cli"
	"github.com/zwoooz/gator/internal/cli/handlers/feeds"
	"github.com/zwoooz/gator/internal/cli/handlers/users"
	"github.com/zwoooz/gator/internal/cli/handlers/util"
	"github.com/zwoooz/gator/internal/config"
	"github.com/zwoooz/gator/internal/database"
	"github.com/zwoooz/gator/internal/rss"
)

func main() {
	os.Exit(run())
}

// run wires everything up and returns the exit code
// os.Exit lives only in main
func run() int {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
		return 1
	}
	if cfg.DBURL == "" {
		fmt.Fprint(os.Stderr, "'db_url' in config is not set\n")
		return 1
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening database: %v\n", err)
		return 1
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error closing database: %v\n", err)
		}
	}()

	rssClient := rss.NewClient(10 * time.Second)

	s := cli.State{
		Cfg:       cfg,
		DB:        database.New(db),
		RssClient: rssClient,
	}

	cmds := cli.NewCommands()
	users.Register(cmds)
	util.Register(cmds)
	feeds.Register(cmds)

	cmd, err := cli.ParseInput(os.Args[1:])
	if err != nil {
		return handleResult(err, cmds, &s, cmd)
	}

	return handleResult(cmds.Run(&s, cmd), cmds, &s, cmd)
}

// handleResult prints err, runs help for usage errors,
// and maps the outcome to an exit code:
// 0 success, 1 failure, 2 incorrect usage
func handleResult(err error, cmds *cli.Commands, s *cli.State, cmd cli.Command) int {
	if err == nil {
		return 0
	}
	fmt.Fprintln(os.Stderr, err)

	if errors.Is(err, cli.ErrUsage) {
		help := cli.Command{Name: "help", Args: []string{cmd.Name}}
		err := cmds.Run(s, help)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		return 2
	}
	return 1
}
