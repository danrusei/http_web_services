package opening

import (
	"log"
)

//OpenRequest holds the request data
type OpenRequest struct {
	ID     int  `json:"id"`
	IsOpen bool `json:"isopen"`
}

// ErrItemNotFound is used when a beer already exists.
//var ErrItemNotFound = errors.New("can't find the item")

// Service provides beer adding operations.
type Service interface {
	OpenItem(OpenRequest) error
}

// Repository provides access to item repository.
type Repository interface {
	OpenItem(OpenRequest) error
}

type service struct {
	r Repository
}

// NewService creates an opening service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// OpenItem flag the given item as being opened
func (s *service) OpenItem(request OpenRequest) error {

	// any validation can be done here

	if err := s.r.OpenItem(request); err != nil {
		log.Printf("could not change the status of the item: %v", err)
		return err

	}
	return nil
}
