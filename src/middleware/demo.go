package middleware

import (
	"net/http"
	"fmt"
	"time"
	"github.com/urfave/negroni"
)

type DemoMiddleware struct {}

func NewDemoMiddleware() *DemoMiddleware {
	return &DemoMiddleware{}
}

func (dmdw *DemoMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	start := time.Now()

	next(w, r)

	res := w.(negroni.ResponseWriter)

	fmt.Printf("\n%v %s %s %v",  res.Status(), r.Method, r.URL.Path, time.Since(start).Seconds())
}

