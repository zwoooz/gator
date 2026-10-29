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

	err = cfg.SetUser("zwoooz")
	if err != nil {
		log.Fatalf("error setting user: %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Config: %+v", cfg)
}
