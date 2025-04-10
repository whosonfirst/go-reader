package reader

import (
	"context"
	_ "io/ioutil"
	_ "os"
	"strings"
	"testing"
)

func TestSchemes(t *testing.T) {

	schemes := Schemes()

	str_schemes := strings.Join(schemes, " ")

	if str_schemes != "cwd:// fs:// null:// repo:// stdin://" {
		t.Fatalf("Unexpected schemes: '%s'", str_schemes)
	}
}

func TestNewReader(t *testing.T) {

	ctx := context.Background()

	schemes := Schemes()

	for _, s := range schemes {

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

		default:
			uri = s
		}

		_, err := NewReader(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to create new reader for %s, %v", uri, err)
		}
	}
}
