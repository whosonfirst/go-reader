package read

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var reader_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("read")

	fs.StringVar(&reader_uri, "reader-uri", "cwd://", "A valid whosonfirst/go-reader/v2 URI")

	return fs
}
