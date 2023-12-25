package services

import "sync"

// знаю, плохая идея, но для тренирвоки дженериков - сойдёт
type Aggregator[TK comparable, T any] interface {
	All() map[TK]T
}

type KeyValueStorage[TKey comparable, TValue any] interface {
	SetValue(TKey, TValue)
	GetValue(TKey) (TValue, bool)
	Aggregator[TKey, TValue]
}

type InMemoryStorage[TKey comparable, TValue any] struct {
	mu   sync.Mutex
	data map[TKey]TValue
}

func (s *InMemoryStorage[K, T]) SetValue(key K, value T) {
	defer s.mu.Unlock()
	s.mu.Lock()
	s.data[key] = value
}

func (s *InMemoryStorage[K, T]) GetValue(key K) (T, bool) {
	defer s.mu.Unlock()
	s.mu.Lock()
	r, ok := s.data[key]
	return r, ok
}

func (s *InMemoryStorage[TKey, TValue]) All() map[TKey]TValue {
	defer s.mu.Unlock()
	userNames := make(map[TKey]TValue, len(s.data))

	s.mu.Lock()
	for k, v := range s.data {
		userNames[k] = v
	}

	return userNames
}
