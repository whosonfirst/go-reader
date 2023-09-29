package reader

import (
	"io"
	"log/slog"
)

// DefaultLogger() returns a `slog.TextHandler` instance that writes to `io.Discard`.
func DefaultLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(io.Discard, nil))
}
