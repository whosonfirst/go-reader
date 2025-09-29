package read

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var reader_uris multi.MultiString
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("read")

	fs.Var(&reader_uris, "reader-uri", "One or more registered whosonfirst/go-reader.Reader URIs.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
