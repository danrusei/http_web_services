package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

var (
	errNotFound = errors.New("Item not found, the operation failed")
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
		valid := checkValidity(item)
		item.IsValid = valid
		listItems = append(listItems, item)
	}

	return listItems, nil
}

func (a *api) addGood(items ...Item) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	for _, item := range items {
		for _, i := range a.db.Items {
			if item.ID == i.ID {
				return "item already exists", nil
			}

		}
		addtime := time.Now().Format(layoutRO)
		addtime1, err := time.Parse(layoutRO, addtime)
		if err != nil {
			log.Printf("Can't parse the date, %v", err)
			return "", fmt.Errorf("can't parse the date: %v", err)
		}

		item.ID = len(a.db.Items) + 1
		item.Created = timestamp{addtime1}

		a.db.Items = append(a.db.Items, item)
	}

	return "Items has been added to database", nil
}

func (a *api) openState(id int, status bool) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	var found int
	for i, item := range a.db.Items {
		if id == item.ID {
			found = i
			break
		}
	}
	if found != 0 {
		opentimeS := time.Now().Format(layoutRO)
		opentimeT, err := time.Parse(layoutRO, opentimeS)
		if err != nil {
			log.Printf("Can't parse the date, %v", err)
			return "", fmt.Errorf("can't parse the date: %v", err)
		}
		a.db.Items[found].IsOpen = status
		a.db.Items[found].Opened = timestamp{opentimeT}

	}
	return fmt.Sprintf("Item with id %d has been opened", found+1), nil
}

func (a *api) delGood(id int) (string, error) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	var index int
	for i, item := range a.db.Items {
		if id == item.ID {
			index = i
			break
		}
	}
	if index != 0 {
		a.db.Items = removeIndex(a.db.Items, index)
		return fmt.Sprintf("Item id %d has been deleted", id), nil
	}
	return "", errNotFound
}

func removeIndex(s []Item, index int) []Item {
	return append(s[:index], s[index+1:]...)
}

func checkValidity(i Item) bool {
	t := time.Now()
	i.IsValid = true
	if t.Sub(i.ExpDate.Time) > 0 {
		i.IsValid = false
	}

	if i.IsOpen {
		//if i.Opened.Time.Add(time.Duration(int64(time.Duration(i.ExpOpen * 24).Hours()))).Before(t) {
		if t.Sub(i.Opened.Time.AddDate(0, 0, i.ExpOpen)) > 0 {
			i.IsValid = false
		}
	}

	return i.IsValid
}
