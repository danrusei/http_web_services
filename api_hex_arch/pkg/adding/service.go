package adding

import (
	"errors"
	"log"
)

// ErrDuplicate is used when a beer already exists.
var ErrDuplicate = errors.New("item already exists")

// Service provides beer adding operations.
type Service interface {
	AddItem(...Item)
	AddSampleItem([]Item)
}

// Repository provides access to item repository.
type Repository interface {
	// AddItem saves a given beer to the repository.
	AddItem(Item) error
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddItem adds the given item(s) to the database
func (s *service) AddItem(items ...Item) {

	// any validation can be done here

	for _, item := range items {
		if err := s.r.AddItem(item); err != nil {
			log.Printf("could not add the item: %v", err)
		}
	}
}

// AddItem adds the given item(s) to the database
func (s *service) AddSampleItem(items []Item) {

	// any validation can be done here

	for _, item := range items {
		if err := s.r.AddItem(item); err != nil {
			log.Printf("could not add the sample data: %v", err)
		}
	}
}
