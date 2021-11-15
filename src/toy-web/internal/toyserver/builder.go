package toyserver

import (
	"fmt"
	"sync"
	tw "toy-web"
)

var (
	serverGenerators = make(map[string]ServerGenerator, 1)
	mutex            sync.Mutex
)

type ServerGenerator func() (tw.IServer, error)

func Register(name string, generator ServerGenerator) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, ok := serverGenerators[name]
	if ok {
		return fmt.Errorf("the server with name [%s] has been registered", name)
	}
	serverGenerators[name] = generator
	return nil
}

func New(name string) (tw.IServer, error) {
	generator, ok := serverGenerators[name]
	if !ok {
		return nil, fmt.Errorf("the server with name [%s] didn't register", name)
	}
	return generator()
}
