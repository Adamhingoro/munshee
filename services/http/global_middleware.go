package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.WithField("Requested URI" ,r.RequestURI).Info("HttpServer")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		r.Header.Set("Sample" , "This is the sample value")
		next.ServeHTTP(w, r)
	})
}