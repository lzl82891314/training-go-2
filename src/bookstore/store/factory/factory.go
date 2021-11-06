package factory

import (
	"bookstore/store"
	"fmt"
	"sync"
)

var (
	factory = make(map[string]store.Store, 5)
	mutex   = sync.RWMutex{}
)

func Register(key string, store store.Store) {
	mutex.Lock()
	defer mutex.Unlock()

	if store == nil {
		panic("the store can not be nil")
	}
	if _, ok := factory[key]; ok {
		panic("the store has existed")
	}
	factory[key] = store
}

func Create(key string) (store.Store, error) {
	if s, ok := factory[key]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("create a store %s should Register firstly", key)
}
