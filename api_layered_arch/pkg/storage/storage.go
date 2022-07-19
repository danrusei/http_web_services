package storage

import "github.com/Danr17/http_web_services/api_layered_arch/pkg/model"

//Storage the behavior for storing and retrieving items
type Storage interface {
	ListsGoods() ([]model.Item, error)
	AddGood(...model.Item) (string, error)
	OpenState(int, bool) (string, error)
	DelGood(int) (string, error)
}
