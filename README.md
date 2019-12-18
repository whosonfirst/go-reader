# go-reader

There are many interfaces for reading files. This one is ours. It returns `io.ReadCloser` instances.

This package supersedes the [go-whosonfirst-readwrite](https://github.com/whosonfirst/go-whosonfirst-readwrite) package.

## Example

Readers are instantiated with the `reader.NewReader` method which takes as its arguments a `context.Context` instance and a URI string. The URI's scheme represents the type of reader it implements and the remaining (URI) properties are used by that reader type to instantiate itself.

For example to read files from a directory on the local filesystem you would write:

```
package main

import (
	"context"
	"github.com/whosonfirst/go-reader"
	"io"
	"os"
)

func main() {
	r, _ := reader.NewReader(ctx, "local:///usr/local/data")
	fh, _ := r.Read(ctx, "example.txt")
	defer fh.Close()
	io.Copy(os.Stdout, fh)
}
```

Note the use of the `local://` scheme rather than the more conventional `file://`. This is deliberate so as not to overlap with the [Go Cloud](https://gocloud.dev/howto/blob/) `Blob` package's file handler.

There is also a handy "null" reader in case you need a "pretend" reader that doesn't actually do anything:

```
package main

import (
	"context"
	"github.com/whosonfirst/go-reader"
	"io"
	"os"
)

func main() {
	r, _ := reader.NewReader(ctx, "null://")
	fh, _ := r.Read(ctx, "example.txt")
	defer fh.Close()
	io.Copy(os.Stdout, fh)
}
```

## Interfaces

### reader.Reader

```
type Reader interface {
	Open(context.Context, string) error
	Read(context.Context, string) (io.ReadCloser, error)
	URI(string) string
}
```

Should this interface have a `Close()` method? Maybe. We'll see.

## Custom readers

Custom readers need to:

1. Implement the interface above.
2. Announce their availability using the `go-reader.Register` method on initialization.

For example, this is how the [go-reader-http](https://github.com/whosonfirst/go-reader-http) reader is implemented:

```
package reader

import (
	"context"
	"errors"
	wof_reader "github.com/whosonfirst/go-reader"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

func init() {
	r := NewHTTPReader()
	wof_reader.Register("http", r)
	wof_reader.Register("https", r)	
}

type HTTPReader struct {
	wof_reader.Reader
	url *url.URL
	throttle <-chan time.Time
}

func NewHTTPReader() wof_reader.Reader {

	rate := time.Second / 3
	throttle := time.Tick(rate)

	r := HTTPReader{
		throttle: throttle,
	}

	return &r
}

func (r *HTTPReader) Open(ctx context.Context, uri string) error {

	u, err := url.Parse(uri)

	if err != nil {
		return err
	}

	r.url = u
	return nil
}

func (r *HTTPReader) Read(ctx context.Context, uri string) (io.ReadCloser, error) {

	<-r.throttle

	u, _ := url.Parse(r.url.String())
	u.Path = filepath.Join(u.Path, uri)
	
	url := u.String()

	rsp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != 200 {
		return nil, errors.New(rsp.Status)
	}

	return rsp.Body, nil
}
```

And then to use it you would do this:

```
package main

import (
	"context"
	"github.com/whosonfirst/go-reader"
	_ "github.com/whosonfirst/go-reader-http"	
	"io"
	"os"
)

func main() {
	r, _ := reader.NewReader(ctx, "https://data.whosonfirst.org")
	fh, _ := r.Read(ctx, "101/736/545/101736545.geojson")
	defer fh.Close()
	io.Copy(os.Stdout, fh)
}
```

## Available readers

### "blob"

* https://github.com/whosonfirst/go-reader-blob

### github://

* https://github.com/whosonfirst/go-reader-github

### githubapi://

* https://github.com/whosonfirst/go-reader-github

### http:// and https://

* https://github.com/whosonfirst/go-reader-http

### local://

* https://github.com/whosonfirst/go-reader

### null://

* https://github.com/whosonfirst/go-reader