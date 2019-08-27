package main

func (a *api) routes() {
	a.router.HandleFunc("/", a.handlerLists())
	a.router.HandleFunc("/add", a.handlerAdd())
	a.router.HandleFunc("/modify", a.handlerModify())
	a.router.HandleFunc("/del", a.handlerDelete())
}
