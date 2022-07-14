package dbmemory

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Danr17/http_web_services/api_layered_arch/pkg/model"
)

var (
	errNotFound = errors.New("Item not found, the operation failed")
)

//Memory save tha data in memory
type Memory struct {
	Items []model.Item
	mutex sync.Mutex
}

//NewMemory instantiate the memory
func NewMemory() *Memory {
	return &Memory{
		mutex: sync.Mutex{},
	}
}

func (m *Memory) ListsGoods() ([]model.Item, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var listItems []model.Item

	for _, item := range m.Items {
		valid := checkValidity(item)
		item.IsValid = valid
		listItems = append(listItems, item)
	}

	return listItems, nil
}

func (m *Memory) AddGood(items ...model.Item) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for _, item := range items {
		for _, i := range m.Items {
			if item.ID == i.ID {
				return "item already exists", nil
			}

		}
		addtime := time.Now().Format(model.LayoutRO)
		addtime1, err := time.Parse(model.LayoutRO, addtime)
		if err != nil {
			log.Printf("Can't parse the date, %v", err)
			return "", fmt.Errorf("can't parse the date: %v", err)
		}

		if len(m.Items) == 0 {
			item.ID = len(m.Items) + 1
		} else {
			lastItem := m.Items[len(m.Items)-1]
			item.ID = lastItem.ID + 1
		}

		item.Created = model.Timestamp{addtime1}

		m.Items = append(m.Items, item)
	}

	return "Items has been added to database", nil
}

func (m *Memory) OpenState(id int, status bool) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var foundIndex int
	var found bool
	for i, item := range m.Items {
		if id == item.ID {
			found = true
			foundIndex = i
			break
		}
	}
	if !found {
		return "", errNotFound
	}

	opentimeS := time.Now().Format(model.LayoutRO)
	opentimeT, err := time.Parse(model.LayoutRO, opentimeS)
	if err != nil {
		log.Printf("Can't parse the date, %v", err)
		return "", fmt.Errorf("can't parse the date: %v", err)
	}
	m.Items[foundIndex].IsOpen = status
	m.Items[foundIndex].Opened = model.Timestamp{opentimeT}

	return fmt.Sprintf("Item with id %d has been opened", m.Items[foundIndex].ID), nil
}

func (m *Memory) DelGood(id int) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var foundIndex int
	var found bool
	for i, item := range m.Items {
		if id == item.ID {
			found = true
			foundIndex = i
			break
		}
	}

	if !found {
		return "", errNotFound
	}

	m.Items = removeIndex(m.Items, foundIndex)
	return fmt.Sprintf("Item id %d has been deleted", id), nil

}

func removeIndex(s []model.Item, index int) []model.Item {
	return append(s[:index], s[index+1:]...)
}

func checkValidity(i model.Item) bool {
	t := time.Now()
	i.IsValid = true
	if t.Sub(i.ExpDate.Time) > 0 {
		i.IsValid = false
	}

	if i.IsOpen {
		if t.Sub(i.Opened.Time.AddDate(0, 0, i.ExpOpen)) > 0 {
			i.IsValid = false
		}
	}

	return i.IsValid
}
