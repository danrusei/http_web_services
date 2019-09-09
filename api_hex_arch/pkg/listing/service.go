package listing

// Repository provides access to the items storage.
type Repository interface {
	// GetGoods returns all the Items.
	ListItems() ([]Item, error)
}

// Service provides beer and review listing operations.
type Service interface {
	ListItems() ([]Item, error)
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetBeers returns all beers
func (s *service) ListItems() ([]Item, error) {
	return s.r.ListItems()
}
