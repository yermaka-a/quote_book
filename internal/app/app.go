package app

import (
	"fmt"
	"log/slog"
	"quote_book/internal/config"
	"quote_book/internal/core/service"
	"quote_book/internal/storage"
	"quote_book/internal/transport/http"
	"quote_book/internal/transport/http/handler"
)

type App struct {
	cfg    config.Config
	logger *slog.Logger
	server *http.Server
}

func New(cfg config.Config, logger *slog.Logger, storage *storage.MemoryQuoteStorage) *App {

	quoteService := service.NewQuoteService(storage, logger)

	quoteHandler := handler.NewQuoteHandler(quoteService, logger)
	server := http.NewServer(quoteHandler)

	return &App{
		cfg:    cfg,
		logger: logger,
		server: server,
	}
}

func (a *App) Run() error {
	a.logger.Info(fmt.Sprintf("starting server on port %s", a.cfg.Port()))
	return a.server.Run(fmt.Sprintf(":%s", a.cfg.Port()))
}
