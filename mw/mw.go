package mw

import "net/http"

// Middleware is a function which will modify or add to a http.Request in some
// way
type Middleware func(next http.Handler) http.Handler

var defaultMW = []Middleware{Logger, Auth}

// Default will add the default middleware stack to the given http.Handler and
// return a handler with the full stack
func Default(next http.HandlerFunc) http.Handler {
	var h http.Handler = http.HandlerFunc(next)
	for _, m := range defaultMW {
		h = m(h)
	}

	return h
}
