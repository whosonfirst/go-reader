package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-whosonfirst-reader"
	"io"
	"log"
	"os"
)

func main() {

	id := flag.Int64("id", -1, "...")
	path := flag.String("path", "", "")

	source := flag.String("source", "", "")

	flag.Parse()

	if *id == -1 && *path == "" {
		log.Println("Missing -id or -path flags.")
	}

	ctx := context.Background()

	r, err := reader.NewReader(ctx, *source)

	if err != nil {
		log.Fatal(err)
	}

	var fh io.ReadCloser

	if *id != -1 {
		fh, err = reader.ReadFromID(ctx, r, *id)
	} else {

		fh, err = r.Read(ctx, *path)
	}

	if err != nil {
		log.Fatal(err)
	}

	defer fh.Close()

	io.Copy(os.Stdout, fh)
}
