package memory

import (
	"fmt"
	"time"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/adding"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/opening"
)

//Storage storage keeps data in memory
type Storage struct {
	items []Item
}

// ListItems return all items
func (m *Storage) ListItems() ([]listing.Item, error) {
	var items []listing.Item

	for i := range m.items {

		valid := checkValidity(m.items[i])

		item := listing.Item{
			ID:      m.items[i].ID,
			Created: m.items[i].Created,
			Good: listing.Good{Name: m.items[i].Name,
				Manufactured: m.items[i].Manufactured,
				ExpDate:      m.items[i].ExpDate,
				ExpOpen:      m.items[i].ExpOpen},
			IsOpen:  m.items[i].IsOpen,
			Opened:  m.items[i].Opened,
			IsValid: valid,
		}

		items = append(items, item)
	}

	return items, nil
}

// AddItem add the item to repository
func (m *Storage) AddItem(it adding.Item) error {
	for _, item := range m.items {
		if it.ID == item.ID {
			return fmt.Errorf("item %d already exists in database", it.ID)
		}
	}

	if len(m.items) == 0 {
		it.ID = len(m.items) + 1
	} else {
		lastItem := m.items[len(m.items)-1]
		it.ID = lastItem.ID + 1
	}

	addtime := time.Now()
	it.Created = addtime

	newItem := Item{
		ID:      it.ID,
		Created: it.Created,
		Good: Good{
			Name:         it.Name,
			Manufactured: it.Manufactured,
			ExpDate:      it.ExpDate,
			ExpOpen:      it.ExpOpen},
		IsOpen:  it.IsOpen,
		Opened:  it.Opened,
		IsValid: it.IsValid,
	}

	m.items = append(m.items, newItem)

	return nil
}

// OpenItem return all items
func (m *Storage) OpenItem(request opening.OpenRequest) error {
	var foundIndex int
	var found bool
	for index, item := range m.items {
		if request.ID == item.ID {
			found = true
			foundIndex = index
			break
		}
	}

	if !found {
		return fmt.Errorf("the item with the id %d couldn't be found", request.ID)
	}

	m.items[foundIndex].IsOpen = request.IsOpen
	m.items[foundIndex].Opened = time.Now()

	return nil
}

// RemoveItem return all items
func (m *Storage) RemoveItem(id int) error {
	var foundIndex int
	var found bool
	for index, item := range m.items {
		if id == item.ID {
			found = true
			foundIndex = index
			break
		}
	}

	if !found {
		return fmt.Errorf("the item with the id %d couldn't be found", id)
	}

	m.items = removeIndex(m.items, foundIndex)

	return nil
}

func checkValidity(i Item) bool {
	t := time.Now()
	i.IsValid = true
	if t.Sub(i.ExpDate) > 0 {
		i.IsValid = false
	}

	if i.IsOpen {
		if t.Sub(i.Opened.AddDate(0, 0, i.ExpOpen)) > 0 {
			i.IsValid = false
		}
	}

	return i.IsValid
}

func removeIndex(s []Item, index int) []Item {
	return append(s[:index], s[index+1:]...)
}
