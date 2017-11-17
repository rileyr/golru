package lru

import "sync"

var _ Cache = &MultiThreaded{}

// MultiThreaded is a mutex wrapper around the Basic
// LRU implementation.
type MultiThreaded struct {
	cache Cache
	lock  *sync.Mutex
}

func newMultiThreaded(c Cache) Cache {
	return &MultiThreaded{
		cache: c,
		lock:  &sync.Mutex{},
	}
}

func (m *MultiThreaded) Add(k, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.cache.Add(k, v)
}

func (m *MultiThreaded) Get(k interface{}) (interface{}, bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.cache.Get(k)
}

func (m *MultiThreaded) Remove(k interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.cache.Remove(k)
}

func (m *MultiThreaded) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache.Clear()
}
