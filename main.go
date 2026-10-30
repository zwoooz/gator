package main

import (
	"fmt"
	"log"

	"github.com/zwoooz/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Config: %+v\n", cfg)

	s := state{
		cfg: &cfg,
	}

	cmd := commands{
		handlers: make(map[string]func(*state, command) error),
	}
}
