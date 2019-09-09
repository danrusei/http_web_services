package removing

import (
	"log"
)

// Service provides beer adding operations.
type Service interface {
	RemoveItem(int) error
}

// Repository provides access to item repository.
type Repository interface {
	RemoveItem(int) error
}

type service struct {
	r Repository
}

// NewService creates an opening service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// OpenItem flag the given item as being opened
func (s *service) RemoveItem(id int) error {

	// any validation can be done here

	if err := s.r.RemoveItem(id); err != nil {
		log.Printf("could not remove the item: %v", err)
		return err

	}
	return nil
}
