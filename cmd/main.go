package main

import (
	"context"
	"log"
	"read-only_reader_service/internal/app"
	"read-only_reader_service/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig()

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	a.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
