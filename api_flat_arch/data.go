package main

import "time"

// PopulateItems insert the items in empty database
func (db *Memory) PopulateItems() {
	defaultItems := []Item{
		{
			ID: 1,
			Good: Good{
				Name:         "Milk",
				Manufactured: time.Date(2019, 7, 23, 00, 00, 00, 00, time.UTC),
				ExpDate:      time.Date(2019, 8, 23, 00, 00, 00, 00, time.UTC),
				ExpOpen:      10,
			},
			IsOpen: false,
			Opened: time.Date(2019, 7, 30, 00, 00, 00, 00, time.UTC),
		},
		{
			ID: 2,
			Good: Good{
				Name:         "Butter",
				Manufactured: time.Date(2019, 7, 15, 00, 00, 00, 00, time.UTC),
				ExpDate:      time.Date(2019, 10, 23, 00, 00, 00, 00, time.UTC),
				ExpOpen:      20,
			},
			IsOpen: false,
			Opened: time.Date(2019, 8, 20, 00, 00, 00, 00, time.UTC),
		},
	}

	for i, item := range defaultItems {

		item.ID = i + 1
		item.Created = time.Now()

		db.Items = append(db.Items, item)
	}

}
