package http

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/kazukousen/gae-go/gae-go/misc"
)

// Chain enables middleware chaining
func Chain(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return chain(true, f)
}

func chain(logging bool, f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return alice.New(timeout).Then(http.HandlerFunc(custom(logging, f)))
}

type customWriter struct {
	io.Writer
	http.ResponseWriter
	status int
}

func (r *customWriter) Write(b []byte) (int, error) {
	if r.Header().Get("Content-Type") == "" {
		r.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return r.Writer.Write(b)
}

func (r *customWriter) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.status = status
}

func custom(logging bool, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		if ip, found := header(r, "X-Forwarded-For"); found {
			addr = ip
		}

		// compress
		ioWriter := w.(io.Writer)
		for _, val := range misc.Split(r.Header.Get("Accept-Encoding"), ",", -1) {
			if val == "gzip" {
				w.Header().Set("Content-Type", "gzip")
				g := gzip.NewWriter(w)
				defer g.Close()
				ioWriter = g
				break
			}
			if val == "deflate" {
				w.Header().Set("Content-Type", "deflate")
				z := zlib.NewWriter(w)
				defer z.Close()
				ioWriter = z
				break
			}
		}

		writer := &customWriter{Writer: ioWriter, ResponseWriter: w, status: http.StatusOK}

		f(writer, r)
		if logging {
			log.Printf("%s, %s, %s", addr, r.Method, r.URL)
		}
	}
}

func header(r *http.Request, key string) (string, bool) {
	if r.Header == nil {
		return "", false
	}
	if candidate := r.Header[key]; !misc.ZeroOrNil(candidate) {
		return candidate[0], true
	}
	return "", false
}

func timeout(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 300*time.Second, "timed out")
}
