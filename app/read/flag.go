package read

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var reader_uri string
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("read")

	fs.StringVar(&reader_uri, "reader-uri", "cwd://", "A registered whosonfirst/go-reader.Reader URI.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	return fs
}
