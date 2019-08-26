package main

import (
	"net/http"
)

// API holds dependencies
type api struct {
	router *http.ServeMux
}

func newAPI() *api {
	a := &api{
		router: http.NewServeMux(),
	}
	a.routes()
	return a
}

func (a *api) routes() {
	a.router.HandleFunc("/", handlerLists)
	a.router.HandleFunc("/add", handlerAdd)
	a.router.HandleFunc("/modify", handlerModify)
	a.router.HandleFunc("/del", handlerDelete)
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
