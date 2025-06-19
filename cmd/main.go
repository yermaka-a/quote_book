package main

import (
	"log"
	"quote_book/internal/app"
	"quote_book/internal/config"
	"quote_book/internal/logger"
	"quote_book/internal/storage"
)

func main() {

	cfg := config.MustLoad()

	logger := logger.Setup()

	storage := storage.New()

	app := app.New(cfg, logger, storage)
	err := app.Run()
	if err != nil {
		log.Fatalf("server running error %+s", err.Error())
	}
}
