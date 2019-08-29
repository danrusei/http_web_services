package main

func (a *api) routes() {
	a.router.HandleFunc("/", a.handleLists())
	a.router.HandleFunc("/add", a.handleAdd())
	a.router.HandleFunc("/modify", a.handleModify())
	a.router.HandleFunc("/del", a.handleDelete())
}
