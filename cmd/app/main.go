package main

import (
	"log"

	"delivery/internal/app"
	"delivery/internal/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
