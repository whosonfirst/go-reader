package reader

import (
	"context"
	"testing"
)

func TestNewReader(t *testing.T) {

	ctx := context.Background()

	for _, s := range ReaderSchemes() {

		var uri string

		switch s {
		case "fs://", "repo://":

			continue

			// Why aren't these being created?

			/*
				path, err := ioutil.TempDir("", "reader")

				if err != nil {
					t.Fatalf("Failed to create temp dir, %v", err)
				}

				defer os.RemoveAll(path)
				uri = s + path
			*/

		case "sql://":

			continue

		default:
			uri = s
		}

		_, err := NewReader(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to create new reader for %s, %v", uri, err)
		}
	}
}
