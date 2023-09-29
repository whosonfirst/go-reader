package reader

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/whosonfirst/go-ioutil"
)

// Constant string value representing STDIN.
const STDIN string = "-"

// StdinReader is a struct that implements the `Reader` interface for reading documents from STDIN.
type StdinReader struct {
	Reader
	logger *slog.Logger
}

func init() {

	ctx := context.Background()
	err := RegisterReader(ctx, "stdin", NewStdinReader)

	if err != nil {
		panic(err)
	}
}

// NewStdinReader returns a new `FileReader` instance for reading documents from STDIN,
// configured by 'uri' in the form of:
//
//	stdin://
//
// Technically 'uri' can also be an empty string.
func NewStdinReader(ctx context.Context, uri string) (Reader, error) {

	logger := DefaultLogger()

	r := &StdinReader{
		logger: logger,
	}

	return r, nil
}

// Read will open a `io.ReadSeekCloser` instance wrapping `os.Stdin`.
func (r *StdinReader) Read(ctx context.Context, uri string) (io.ReadSeekCloser, error) {
	return ioutil.NewReadSeekCloser(os.Stdin)
}

// ReaderURI will return the value of the `STDIN` constant.
func (r *StdinReader) ReaderURI(ctx context.Context, uri string) string {
	return STDIN
}

// SetLogger assigns 'logger' to 'r'.
func (r *StdinReader) SetLogger(ctx context.Context, logger *slog.Logger) error {
	r.logger = logger
	return nil
}
