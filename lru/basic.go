package lru

import (
	"container/list"
)

var _ Cache = &Basic{}

// Basic is a basic LRU implementation.
type Basic struct {
	list  *list.List
	store map[interface{}]*list.Element
	size  int
}

type item struct {
	key   interface{}
	value interface{}
}

func newBasic(size int) Cache {
	return &Basic{
		list:  list.New(),
		size:  size,
		store: make(map[interface{}]*list.Element, size),
	}
}

func (b *Basic) Add(k, v interface{}) (ok bool) {
	elem, found := b.store[k]
	if found {
		i := elem.Value.(*item)
		i.value = v
		b.list.MoveToFront(elem)
		return false
	}

	if len(b.store) >= b.size {
		b.evictOne()
		ok = true
	}

	i := &item{key: k, value: v}
	elem = b.list.PushFront(i)
	b.store[k] = elem
	return ok
}

func (b *Basic) Get(k interface{}) (v interface{}, ok bool) {
	elem, found := b.store[k]
	if !found {
		return v, ok
	}

	ok = true
	b.list.MoveToFront(elem)
	i := elem.Value.(*item)
	return i.value, ok
}

func (b *Basic) Remove(k interface{}) (ok bool) {
	elem, found := b.store[k]
	if !found {
		return ok
	}

	ok = true
	b.remove(elem)
	return ok
}

func (b *Basic) Clear() {
	for _, v := range b.store {
		b.remove(v)
	}
}

func (b *Basic) evictOne() {
	if len(b.store) == 0 {
		return
	}

	e := b.list.Back()
	b.remove(e)
}

func (b *Basic) remove(e *list.Element) {
	i := e.Value.(*item)
	delete(b.store, i.key)
	b.list.Remove(e)
}
