package cli

import (
	"github.com/zwoooz/gator/internal/config"
	"github.com/zwoooz/gator/internal/database"
)

type State struct {
	Cfg *config.Config
	DB  *database.Queries
}
