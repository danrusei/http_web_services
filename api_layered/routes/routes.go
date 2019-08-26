package routes

import "net/http"

type api struct {
	router *http.ServeMux
}

//NewAPI is api struct constructor
func NewAPI() *api {
	a := &api{}
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
