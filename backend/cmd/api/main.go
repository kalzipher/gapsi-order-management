package main

import (
	"log"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/app"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/config"
)

func main() {
	cfg := config.Load()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("could not initialize app: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
