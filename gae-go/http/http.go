package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/kazukousen/gae-go/gae-go/misc"
)

// Chain enables middleware chaining
func Chain(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return chain(true, f)
}

func chain(log bool, f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return alice.New(timeout).Then(http.HandlerFunc(custom(log, f)))
}

func custom(log bool, f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		addr := r.RemoteAddr
		if ip, found := header(r, "X-Forwarded-For"); found {
			addr = ip
		}
		f(w, r)
		if log {
			fmt.Printf("%s, %s, %s", addr, r.Method, r.URL)
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
