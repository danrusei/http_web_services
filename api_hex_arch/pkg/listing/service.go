package listing

// Repository provides access to the items storage.
type Repository interface {
	// GetGoods returns all the Items.
	ListGoods() ([]Item, error)
}

// Service provides beer and review listing operations.
type Service interface {
	ListGoods() ([]Item, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) ListGoods() ([]Item, error) {
	return s.r.ListGoods()
}
