package memory

import (
	"fmt"
	"time"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/adding"
	"github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"
)

//Storage storage keeps data in memory
type Storage struct {
	items []Item
}

// ListGoods return all items
func (m *Storage) ListGoods() ([]listing.Item, error) {
	var items []listing.Item

	for i := range m.items {

		item := listing.Item{
			ID:      m.items[i].ID,
			Created: m.items[i].Created,
			Good: listing.Good{Name: m.items[i].Name,
				Manufactured: m.items[i].Manufactured,
				ExpDate:      m.items[i].ExpDate,
				ExpOpen:      m.items[i].ExpOpen},
			IsOpen:  m.items[i].IsOpen,
			Opened:  m.items[i].Opened,
			IsValid: m.items[i].IsValid,
		}

		items = append(items, item)
	}

	return items, nil
}

// AddItem add the item to repository
func (m *Storage) AddItem(it adding.Item) error {
	for _, i := range m.items {
		if it.ID == i.ID {
			return fmt.Errorf("item %d already exists in database", it.ID)
		}

		addtime := time.Now()
		it.ID = len(m.items) + 1
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
	}

	return nil
}
