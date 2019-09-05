package main

import (
	"log"
	"net/http"
	"time"
)

func (a *api) logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			log.Printf("The client %s requested %v \n", r.RemoteAddr, r.URL)
			log.Printf("It took %s to serve the request \n", time.Now().Sub(startTime))
		}()
		h(w, r)
	}
}
