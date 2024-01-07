package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Match the key that we want
			if a.Key == slog.TimeKey {
				a.Key = "date" // rename to date
				a.Value = slog.Int64Value(time.Now().Unix())
			}
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.Int("What's the meaning of life?", 42),
		slog.Group("votes",
			slog.Int("Pikachu", 40),
			slog.Int("Mew", 24),
		),
	})

	//logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	//	Level:     slog.LevelDebug,
	//	AddSource: true,
	//})

	logger := slog.New(logHandler)

	logger.Debug("Best Pokemon Rating")

	slog.SetDefault(logger)

	slog.Info("New Info")

	//logger.Debug("What's the meaning of lift?", slog.Int("answer", 42))

	//logger.Debug("debug level")
	//logger.Info("info level")
	//logger.Warn("warn level")
	//logger.Error("error level")
}
