# gator

gator is a command-line RSS aggregator written in Go as a guided project for boot.dev's backend developer path. It is designed to be a long-running aggregator: it continually fetches the feeds you follow, stores their posts in a PostgreSQL database, and lets you read the latest ones from your terminal.

## Features

* Multiple users, with the current user stored in a local config file
* Add feeds by name and URL; users can follow or unfollow feeds added by anyone
* A long-running aggregator that fetches feeds on an interval you choose
* Posts are stored in PostgreSQL, duplicates are skipped automatically
* Browse the latest posts from the feeds you follow

## Getting Started

### Dependencies

* Go (>= 1.27.3), needed to build or install gator
* PostgreSQL, running and reachable, needed whenever gator runs
* Git, needed to fetch the migration files
* [goose](https://github.com/pressly/goose), used once to set up the database schema

Optional:

* [sqlc](https://sqlc.dev) (only if you want to change the SQL)

### Installation

#### With `go install`

```bash
$ go install github.com/zwoooz/gator@latest
```

This places the `gator` binary in `~/go/bin`, so make sure that directory is on your `PATH`.

#### From source

```bash
$ git clone https://github.com/zwoooz/gator.git
$ cd gator
$ go build
```

This creates a `gator` binary in the project directory. The examples below use `gator`; if you built from source and the binary isn't on your `PATH`, run it as `./gator`.

Note that the database setup below needs the repository's migration files, so even if you used `go install` you will need to clone the repository once.

### Initial setup

1. Create a `.gatorconfig.json` in your home directory with the following contents:

```json
{
    "db_url": "postgres://USER:PASSWORD@localhost:5432/gator?sslmode=disable"
}
```

`USER` and `PASSWORD` are the credentials for your PostgreSQL server, and `5432` is the default port. Change it if your database listens elsewhere.

If the file doesn't exist, gator creates one with empty values the first time you run it and exits. Fill in your `db_url` and run it again. The `login` and `register` commands store the current user in this file, so you don't need to edit that part by hand. Do note this file contains your database password.

2. Create the database:

```bash
$ createdb gator
```

3. Run the migrations. Clone the repository if you haven't already (skip the first two lines if you built from source), then install [goose](https://github.com/pressly/goose) and apply the migrations. Use the same URL as in your config:

```bash
$ git clone https://github.com/zwoooz/gator.git
$ cd gator
$ go install github.com/pressly/goose/v3/cmd/goose@latest
$ goose -dir sql/schema postgres "postgres://USER:PASSWORD@localhost:5432/gator?sslmode=disable" up
```

`up` applies every migration, so a fresh database ends up on the latest version. To check, replace `up` with `status`. Like gator, goose is installed to `~/go/bin`.

## Usage

Run `gator help` to see every command and its usage:

```text
Usage:
	gator <COMMAND> [ARGUMENTS]

Util:
	help | help [command]    Shows usage information

Users:
	login [username]         Sets the current active user
	register [username]      Creates a new user and logs in as them
	reset                    Wipes every row on every table in the database (asks for confirmation)
	users                    Lists every user in the database, shows who is currently logged in

Feeds:
	addfeed [name] [URL]     Adds feed to the current logged in user
	agg [interval]           Fetches the current least updated feed on the given interval (e.g. 30s, 5m or 1h), minimum 10s, keeps running until stopped with Ctrl+C
	browse | browse [limit]  Shows saved posts, with optional limit on amount of posts
	feeds                    List all tracked feeds and who created them
	follow [URL]             Follows an already tracked feed as the logged in user
	following                List all feeds followed by the logged in user
	unfollow [URL]           Unfollows a followed feed as the logged in user
```

> **Warning:** `reset` deletes all data in every table. It asks for confirmation first, but it can't be undone.

### Example

```bash
$ gator register zwoooz
$ gator addfeed "Wagslane Blog" https://www.wagslane.dev/index.xml
$ gator agg 1m      # leave this running in its own terminal
$ gator browse 10
```

`agg` runs until you stop it with Ctrl+C, so use a second terminal for `browse` while it runs.
