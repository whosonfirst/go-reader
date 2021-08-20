package reader

import (
	"bytes"
	"context"
	_ "fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestStdinReader(t *testing.T) {
	data := []byte("hello world")
	restore := mockStdin(data)
	defer restore()

	ctx := context.Background()

	r, err := NewReader(ctx, "stdin://")

	if err != nil {
		t.Fatal(err)
	}

	fh, err := r.Read(ctx, "-")

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	res, err := io.ReadAll(fh)

	if err != nil {
		t.Fatalf("Failed to read file handle, %v", err)
	}

	if bytes.Compare(res, data) != 0 {
		t.Fatalf("Failed to read file contents, got '%s', expecting '%s'", res, data)
	}
}

// mock stdin for testing
func mockStdin(data []byte) func() {

	// create a new temp file
	tmpfile, err := ioutil.TempFile("", "mock-stdin")
	if err != nil {
		log.Fatal(err)
	}

	// write data
	if _, err := tmpfile.Write(data); err != nil {
		log.Fatal(err)
	}

	// seek to beginning
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	// save a copy of the 'original stdin'
	originalStdin := os.Stdin

	// set os.Stdin to the temporary file
	os.Stdin = tmpfile

	// return a 'restore' function which removes the mock
	return func() {
		// restore original stdin
		os.Stdin = originalStdin

		// remove temp file
		os.Remove(tmpfile.Name())
	}
}
