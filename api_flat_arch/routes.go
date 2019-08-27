package main

import (
	"net/http"
)

// API holds dependencies
type api struct {
	db     *Memory
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
	a.router.HandleFunc("/", a.handlerLists())
	a.router.HandleFunc("/add", a.handlerAdd())
	a.router.HandleFunc("/modify", a.handlerModify())
	a.router.HandleFunc("/del", a.handlerDelete())
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
