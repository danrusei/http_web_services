package handlers

func (s *server) routes() {
	s.router.HandleFunc("/add/", s.handleAddItem())
	s.router.HandleFunc("/remove/", s.handleRemoveItem())
	s.router.HandleFunc("/remove/", s.handleModoifyuItem())
	s.router.HandleFunc("/list", s.handleListItems())
	s.router.HandleFunc("/", s.Logger(s.handleIndex()))
}
