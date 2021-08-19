package reader

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/whosonfirst/go-ioutil"
	"io"
	"os"
)

type StdinReader struct {
	Reader
}

func init() {

	ctx := context.Background()
	err := RegisterReader(ctx, "stdin", NewStdinReader)

	if err != nil {
		panic(err)
	}
}

func NewStdinReader(ctx context.Context, uri string) (Reader, error) {

	r := &StdinReader{}
	return r, nil
}

func (r *StdinReader) Read(ctx context.Context, uri string) (io.ReadSeekCloser, error) {

	var b bytes.Buffer
	wr := bufio.NewWriter(&b)

	_, err := io.Copy(wr, os.Stdin)

	if err != nil {
		return nil, err
	}

	wr.Flush()

	br := bytes.NewReader(b.Bytes())
	return ioutil.NewReadSeekCloser(br)
}

func (r *StdinReader) ReaderURI(ctx context.Context, uri string) string {
	return fmt.Sprintf("stdin://%s", uri)
}
