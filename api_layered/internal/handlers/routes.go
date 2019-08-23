package handlers

func (s *Server) routes() {
	s.router.HandleFunc("/add/", s.handleAddItem)
	s.router.HandleFunc("/remove/", s.handleRemoveItem)
	s.router.HandleFunc("/remove/", s.handleModifyItem)
	s.router.HandleFunc("/", s.log(s.handleListItems))
}
