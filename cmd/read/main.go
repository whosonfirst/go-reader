package main

import (
	"context"
	"log/slog"
	"os"
	
	"github.com/whosonfirst/go-reader/v2/app/read"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := read.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to create new reader, %v", err)
		os.Exit(1)
	}
}
