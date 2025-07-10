package read

import (
	"context"
	"flag"
	"fmt"
	"io"
	_ "log/slog"
	"os"

	"github.com/whosonfirst/go-reader/v2"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if opts.Verbose {

	}

	r, err := reader.NewReader(ctx, opts.ReaderURI)

	if err != nil {
		return fmt.Errorf("Failed to create new reader, %w", err)
	}

	for _, path := range opts.Paths {

		r, err := r.Read(ctx, path)

		if err != nil {
			return fmt.Errorf("Failed to read '%s', %w", path, err)
		}

		defer r.Close()

		_, err = io.Copy(os.Stdout, r)

		if err != nil {
			return fmt.Errorf("Failed to copy contents of '%s' to STDOUT, %w", path, err)
		}
	}

	return nil
}
