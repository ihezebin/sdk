package mgoc

import (
	"github.com/pkg/errors"
	"sync"
)

type Manager interface {
	Add(alias string, c Client)
	Delete(alias string)
	Get(alias string) (Client, error)
	Has(alias string) bool
}

var gManager = manager{new(sync.Map)}

func ClientManager() Manager {
	return &gManager
}

// manager It is used to manage and initialize multiple Mongo clients.
// A service may use multiple Mongo clients. Centralized management is better
type manager struct {
	clientMap *sync.Map
}

// Add In order to avoid the overwriting problem caused by adding the client with the same alias,
// it is recommended to use Has to determine whether the client with the alias exists.
// 为了避免添加相同别名的客户端时导致的覆盖问题，建议先使用Has判断该别名的客户端是否存在
func (manager *manager) Add(alias string, c Client) {
	manager.clientMap.Store(alias, c)
}

func (manager *manager) Delete(alias string) {
	manager.clientMap.Delete(alias)
}

// Get get client by alias, if not exist that will return error
func (manager *manager) Get(alias string) (Client, error) {
	c, ok := manager.clientMap.Load(alias)
	if !ok {
		return nil, errors.Errorf("can't find mongoc client of alias '%s'", alias)
	}
	return c.(Client), nil
}

func (manager *manager) Has(alias string) bool {
	_, ok := manager.clientMap.Load(alias)
	return ok
}
