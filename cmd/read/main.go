package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-reader"
	"io"
	"log"
	"os"
)

func main() {

	reader_uri := flag.String("reader-uri", "file://", "")

	flag.Parse()
	
	uris := flag.Args()
	
	ctx := context.Background()

	r, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatalf("Failed to create new reader, %v", err)
	}

	for _, path := range uris {
		
		fh, err := r.Read(ctx, path)
		
		if err != nil {
			log.Fatalf("Failed to read '%s', %v", path, err)
		}
		
		defer fh.Close()
		
		_, err = io.Copy(os.Stdout, fh)

		if err != nil {
			log.Fatalf("Failed to copy contents of '%s' to STDOUT, %v", path, err)
		}
	}
		
}
