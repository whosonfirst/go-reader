package read

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	ReaderURIs []string   `json:"reader_uri"`
	Paths     []string `json:"paths"`
	Verbose   bool     `json:"verbose"`
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "WHOSONFIRST")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	paths := fs.Args()

	opts := &RunOptions{
		ReaderURIs: reader_uris,
		Paths:     paths,
		Verbose:   verbose,
	}

	return opts, nil
}
