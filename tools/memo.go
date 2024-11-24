package tools

import (
	"log"
	"sync"
)

type Value string
type KeyValue map[string]string

type memory struct {
	storage    KeyValue
	controller sync.RWMutex
}

func (m *memory) Log() {
	go log.Println(m.storage)
}

func (m *memory) SET(key string, value string) {
	defer m.Log()

	m.controller.Lock()
	m.storage[key] = value
	m.controller.Unlock()
}

func (m *memory) GET(key string) (result string) {
	defer m.Log()

	m.controller.Lock()
	result = string(m.storage[key])
	m.controller.Unlock()
	return result
}

var Memory memory = memory{storage: KeyValue{}}
