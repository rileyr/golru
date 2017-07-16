package golru

import "sync"

var _ Cache = &MultiThreaded{}

// MultiThreaded is a mutex wrapper around the Basic
// LRU implementation.
type MultiThreaded struct {
	cache *Basic
	lock  *sync.RWMutex
}

func newMultiThreaded(size int) *MultiThreaded {
	return &MultiThreaded{
		cache: newBasic(size),
		lock:  &sync.RWMutex{},
	}
}

func (m *MultiThreaded) Add(k, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.cache.Add(k, v)
}

func (m *MultiThreaded) Get(k interface{}) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.cache.Get(k)
}

func (m *MultiThreaded) Remove(k interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.cache.Remove(k)
}
