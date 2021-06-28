// package www provides common net/http handlers for reader-related HTTP requests.
package www

import (
	"net/http"
	"io"
	"github.com/whosonfirst/go-reader"	
	"github.com/whosonfirst/go-whosonfirst-uri"
	"strconv"
	_ "log"
)

// The name of the HTTP response header where the Who's On First relative path
// derived from a request URI will be stored.
const HEADER_RELPATH string = "X-WhosOnFirst-Rel-Path"

// Parse a request URI (or ?id query string) in to a valid Who's On First ID relative
// path and assign the value to the 'X-WhosOnFirst-Rel-Path' response header. If 'next'
// is not nil then delegate to that handler, otherwise print the relative path to the
// response handler.
func ParseURIHandler(next http.Handler) http.HandlerFunc {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		path := req.URL.Path
		
		wofid, uri_args, err := uri.ParseURI(path)
		
		if err != nil || wofid == -1 {
			
			q := req.URL.Query()
			str_id := q.Get("id")
			
			if str_id == "" {
				http.Error(rsp, err.Error(), http.StatusNotFound)
				return
			}
			
			id, err := strconv.ParseInt(str_id, 10, 64)
			
			if err != nil {
				http.Error(rsp, err.Error(), http.StatusBadRequest)
				return
			}
			
			wofid = id
			
			uri_args = &uri.URIArgs{
				IsAlternate: false,
			}
		}

		rel_path, err := uri.Id2RelPath(wofid, uri_args)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}
		
		rsp.Header().Set(HEADER_RELPATH, rel_path)
		
		if next != nil {
			next.ServeHTTP(rsp, req)
			return
		}

		rsp.Write([]byte(rel_path))
		return	
	}

	return http.HandlerFunc(fn)
}

// Emit an HTTP redirect header for the value of r.ReaderURI for a URI. This handler is meant to be used
// in conjunction with ParseURIHandler middleware handler. For example:
// 	handler, _ := www.RedirectHandler(r)
//	handler = www.ParseURIHandler(handler)
func RedirectHandler(r reader.Reader) (http.HandlerFunc, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()

		rel_path := rsp.Header().Get(HEADER_RELPATH)		

		if rel_path == "" {
			http.Error(rsp, "Unable to determine URI", http.StatusNotFound)
			return			
		}

		abs_path := r.ReaderURI(ctx, rel_path)
		http.Redirect(rsp, req, abs_path, http.StatusFound)
		return
	}

	h := http.HandlerFunc(fn)
	return h, nil	
}

// Emit the out of r.Read for a URI. This handler is meant to be used
// in conjunction with ParseURIHandler middleware handler. For example:
// 	handler, _ := www.DataHandler(r)
//	handler = www.ParseURIHandler(handler)
func DataHandler(r reader.Reader) (http.HandlerFunc, error) {

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		ctx := req.Context()
		
		rel_path := rsp.Header().Get(HEADER_RELPATH)
		
		if rel_path == "" {
			http.Error(rsp, "Unable to determine URI", http.StatusNotFound)			
			return			
		}

		fh, err := r.Read(ctx, rel_path)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		
		_, err = io.Copy(rsp, fh)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	h := http.HandlerFunc(fn)
	return h, nil	
}
