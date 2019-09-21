package handlers

import (
	"net/http"
	"time"
)

func (h *Handlers) logging(hnext http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer func() {
			h.logger.Printf("The client %s requested %v \n", r.RemoteAddr, r.URL)
			h.logger.Printf("It took %s to serve the request \n", time.Now().Sub(startTime))
		}()
		hnext(w, r)
	}
}
