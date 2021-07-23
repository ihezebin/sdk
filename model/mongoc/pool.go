package mongoc

import (
	"github.com/pkg/errors"
	"sync"
)

// clientMap It is used to manage and initialize multiple Mongo clients.
// A service may use multiple Mongo clients. Centralized management is better
var clientMap = sync.Map{}

func AddClient(alias string, c Client) {
	clientMap.Store(alias, c)
}

func DeleteClient(alias string) {
	clientMap.Delete(alias)
}

func GetClient(alias string) Client {
	c, ok := clientMap.Load(alias)
	if !ok {
		panic(errors.Errorf("can't find mongo client of alias '%s'", alias))
	}
	return c.(Client)
}

func HasClient(alias string) bool {
	_, ok := clientMap.Load(alias)
	return ok
}
