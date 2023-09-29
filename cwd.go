package reader

import (
	"context"
	"fmt"
	"net/url"
	"os"
)

func init() {

	ctx := context.Background()

	schemes := []string{
		"cwd",
	}

	for _, scheme := range schemes {

		err := RegisterReader(ctx, scheme, NewCwdReader)

		if err != nil {
			panic(err)
		}
	}
}

// NewCwdReader returns a new `CwdReader` instance for writing documents to the current working directory
// configured by 'uri' in the form of:
//
//	cwd://
//
// Technically 'uri' can also be an empty string.
func NewCwdReader(ctx context.Context, uri string) (Reader, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()

	cwd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("Failed to derive current working directory, %w", err)
	}

	fs_uri := url.URL{}
	fs_uri.Scheme = "fs"
	fs_uri.Path = cwd
	fs_uri.RawQuery = q.Encode()

	return NewFileReader(ctx, fs_uri.String())
}
