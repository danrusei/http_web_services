package memory

import "github.com/Danr17/http_web_services/api_hex_arch/pkg/listing"

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
