package read

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-reader/v2"
)

func Run(ctx context.Context, logger *slog.Logger) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	flagset.Parse(fs)

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		return fmt.Errorf("Failed to create new reader, %w", err)
	}

	err = r.SetLogger(ctx, logger)

	if err != nil {
		return fmt.Errorf("Failed to set logger, %w", err)
	}

	for _, path := range fs.Args() {

		fh, err := r.Read(ctx, path)

		if err != nil {
			return fmt.Errorf("Failed to read '%s', %v", path, err)
		}

		defer fh.Close()

		_, err = io.Copy(os.Stdout, fh)

		if err != nil {
			return fmt.Errorf("Failed to copy contents of '%s' to STDOUT, %v", path, err)
		}
	}

	return nil
}
