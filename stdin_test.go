package reader

import (
	"context"
	"io"
	_ "fmt"
	"testing"
)

func TestStdinReader(t *testing.T) {

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

	_, err = io.ReadAll(fh)

	if err != nil {
		t.Fatalf("Failed to read filehandle, %v", err)
	}
}
