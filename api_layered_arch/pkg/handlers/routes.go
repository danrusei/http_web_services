package handlers

import "net/http"

//CreateRoutes wire up the server endpoints
func (h *Handlers) CreateRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.logging(h.handleLists()))
	mux.HandleFunc("/add", h.logging(h.handleAdd()))
	mux.HandleFunc("/open", h.logging(h.handleOpen()))
	mux.HandleFunc("/del", h.logging(h.handleDelete()))
}
