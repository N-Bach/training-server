package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
	"fmt"
)

type LoggingMiddleware struct{}

func NewLoggingMiddleware() *LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (l *LoggingMiddleware) ServeHTTP(
	rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	next(rw, r)
	res := rw.(negroni.ResponseWriter)
	fmt.Printf("%d %s %s \n", res.Status(), r.Method, r.URL.Path)
}
