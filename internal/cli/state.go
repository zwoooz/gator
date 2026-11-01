package cli

import (
	"github.com/zwoooz/gator/internal/config"
	"github.com/zwoooz/gator/internal/database"
	"github.com/zwoooz/gator/internal/rss"
)

// State holds the application dependencies shared by all command handlers
type State struct {
	Cfg       *config.Config
	DB        *database.Queries
	RssClient *rss.Client
}
