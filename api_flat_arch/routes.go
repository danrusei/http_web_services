package main

func (a *api) routes() {
	a.router.HandleFunc("/", a.handleLists())
	a.router.HandleFunc("/add", a.handleAdd())
	a.router.HandleFunc("/open", a.handleOpen())
	a.router.HandleFunc("/del", a.handleDelete())
}
