package handlers

import (
	"log"
	"net/http"
)

//Server struct define dependencies
type Server struct {
	db     map[string]interface{}
	logger *log.Logger
	router *http.ServeMux
}

// newServer constructs the server to handle a set of routes.
func newServer() *Server {
	s := &Server{}
	s.routes()
	return s
}

//make server an http.Handle, to pass execution to the router
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleAddItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Item Added"))
}

func (s *Server) handleRemoveItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Item Removed"))
}

func (s *Server) handleModifyItem(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Item Modified"))
}

func (s *Server) handleListItems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Items Listed"))
}
