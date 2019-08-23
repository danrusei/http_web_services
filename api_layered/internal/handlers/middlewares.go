package handlers

import (
	"net/http"
	"time"
)

func (s *Server) log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer s.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
