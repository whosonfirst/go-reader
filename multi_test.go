package reader

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestMultiReader(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatalf("Failed to determine current working directory, %v", err)
	}

	uris := []string{
		fmt.Sprintf("fs://%s/fixtures", cwd),
		"null://",
	}

	readers := make([]Reader, len(uris))

	for idx, u := range uris {

		r, err := NewReader(ctx, u)

		if err != nil {
			t.Fatalf("Failed to create reader for '%s', %v", u, err)
		}

		readers[idx] = r
	}

	mr, err := NewMultiReader(ctx, readers...)

	if err != nil {
		t.Fatalf("Failed to create new multi reader, %v", err)
	}

	to_test := []string{
		"101/736/545/101736545.geojson",
		"101/736/545/101736545.geojson.bz2",
		"do-not-exist.txt", // should be skipped by fs:// reader and accepted by null:// reader
	}

	for _, path := range to_test {

		fh, err := mr.Read(ctx, path)

		if err != nil {
			t.Fatalf("Failed to read %s, %v", path, err)
		}

		defer fh.Close()

		_, err = io.Copy(ioutil.Discard, fh)

		if err != nil {
			t.Fatalf("Failed to copy %s, %v", path, err)
		}
	}
}
