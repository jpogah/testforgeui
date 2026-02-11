package main

import (
	"log"

	"github.com/jpogah/testforgeui/internal/config"
	"github.com/jpogah/testforgeui/internal/webservice"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server := webservice.NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatalf("server exited with error: %v", err)
	}
}
