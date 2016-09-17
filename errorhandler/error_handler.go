package errorhandler

import (
	"log"
	"net/http"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		log.Printf("Got error while processing the request: %s\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
