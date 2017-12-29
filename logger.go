package main

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Logger wraps the handler function with basic log message.
func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		log.Printf("[%s %s] Requested", r.Method, r.URL.Path)
		start := time.Now()
		fn(w, r, param)
		log.Printf("[%s %s] Completed in %v", r.Method, r.URL.Path, time.Since(start))
	}
}
