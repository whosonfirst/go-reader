package reader

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

func TestCwdReader(t *testing.T) {

	ctx := context.Background()

	source := "cwd://?allow_bz2=1"

	r, err := NewReader(ctx, source)

	if err != nil {
		t.Fatal(err)
	}

	to_test := []string{
		"fixtures/101/736/545/101736545.geojson",
		"fixtures/101/736/545/101736545.geojson.bz2",
	}

	for _, path := range to_test {

		fmt.Printf("Read %s\n", path)

		fh, err := r.Read(ctx, path)

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
