package adding

import "errors"

// ErrDuplicate is used when a beer already exists.
var ErrDuplicate = errors.New("item already exists")

// Service provides beer adding operations.
type Service interface {
	AddItem(...Item)
}

// Repository provides access to item repository.
type Repository interface {
	// AddItem saves a given beer to the repository.
	AddItem(Item) error
}

type service struct {
	iR Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddItem adds the given item(s) to the database
func (s *service) AddItem(items ...Item) {

	// any validation can be done here

	for _, item := range items {
		_ = s.iR.AddItem(item) // error handling omitted for simplicity
	}
}
