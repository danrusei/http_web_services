package main

func (a *api) routes() {
	a.router.HandleFunc("/", a.logger(a.handleLists()))
	a.router.HandleFunc("/add", a.logger(a.handleAdd()))
	a.router.HandleFunc("/open", a.logger(a.handleOpen()))
	a.router.HandleFunc("/del", a.logger(a.handleDelete()))
}
