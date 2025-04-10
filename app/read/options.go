package read

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	ReaderURI string   `json:"reader_uri"`
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
		ReaderURI: reader_uri,
		Paths:     paths,
		Verbose:   verbose,
	}

	return opts, nil
}
