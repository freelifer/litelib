package module

import (
	"fmt"
)

type Module interface {
	Setup(config map[string]string) error
}

// Instance is a function create a new Cache Instance
type Instance func() Module

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("module: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("module: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewModule(adapterName string) (adapter Module, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("module: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	return
}
