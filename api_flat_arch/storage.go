package main

import (
	"time"
)

//Memory save tha data in memory
type Memory struct {
	Items []Item
}

func (a *api) listsGoods() ([]Item, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	var listItems []Item

	for _, item := range a.db.Items {
		listItems = append(listItems, item)
	}

	return listItems, nil
}

func (a *api) addGood(items ...Item) (string, error) {
	for _, item := range items {
		for _, i := range a.db.Items {
			if item.ID == i.ID {
				return "item already exists", nil
			}

		}

		item.ID = len(a.db.Items) + 1
		item.Created = time.Now()

		a.db.Items = append(a.db.Items, item)
	}

	return "Items added to database", nil
}

func (a *api) modifyGood(i Item) (string, error) {
	return "", nil
}

func (a *api) delGood(i Item) (string, error) {
	return "", nil
}
