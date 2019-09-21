package storage

import "github.com/Danr17/items-rest-api/pkg/model"

//Storage the behavior for storing and retrieving items
type Storage interface {
	ListsGoods() ([]model.Item, error)
	AddGood(...model.Item) (string, error)
	OpenState(int, bool) (string, error)
	DelGood(int) (string, error)
}
