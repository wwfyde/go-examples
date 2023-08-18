package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	// Create a new logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	// logging a structured info
	logger.Info("my first zap",
		zap.Int("attempt", 3),
		zap.Duration("duration", time.Second),
	)

}
