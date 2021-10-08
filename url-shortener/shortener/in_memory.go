package shortener

import (
	"sync"
)

type InMemory struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewInMemory() *InMemory {
	return &InMemory{urls: make(map[string]string)}
}

func (s *InMemory) Get(id string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.urls[id]
}

func (s *InMemory) Put(value string) string {
	id := generateID(value)

	s.mu.Lock()
	defer s.mu.Unlock()
	if existing, ok := s.urls[id]; ok {
		if value == existing {
			return id
		}
		panic("id collision detected")
	}
	s.urls[id] = value
	return id
}
