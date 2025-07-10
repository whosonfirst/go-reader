package reader

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"testing"
)

func TestHTTPReader(t *testing.T) {

	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:10.0) Gecko/20100101 Firefox/10.0"

	reader_uri := fmt.Sprintf("https://data.whosonfirst.org?user-agent=%s", url.QueryEscape(ua))
	file_uri := "101/736/545/101736545.geojson"

	ctx := context.Background()

	r, err := NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatal(err)
	}

	fh, err := r.Read(ctx, file_uri)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatal(err)
	}

}
