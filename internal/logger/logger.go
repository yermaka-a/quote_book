package logger

import (
	"log/slog"
	"os"
)

func Setup() *slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			}),
	)
	return logger
}

