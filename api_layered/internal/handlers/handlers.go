package handlers

import (
	"log"
	"net/http"
)

type server struct {
	logger *log.Logger
	router *http.ServeMux
}

// newServer constructs the server to handle a set of routes.
func newServer() *server {
	s := &server{}
	s.routes()
	return s
}

//make server an http.Handle, to pass execution to the router
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
