package golru

var _ Cache = &LazyLookup{}

// LazyLookup is a cache that will
// attempt to look up and cache any Gets that miss.
type LazyLookup struct {
	cache    Cache
	lookupFn LookupFunc
}

// LookupFunc is any function that can provide the value for a given
// key. If the value cannot be obtained, false should be returned.
type LookupFunc func(interface{}) (interface{}, bool)

func newLazyLookup(c Cache, fn LookupFunc) Cache {
	return &LazyLookup{cache: c, lookupFn: fn}
}

func (l *LazyLookup) Add(k, v interface{}) bool {
	return l.cache.Add(k, v)
}

func (l *LazyLookup) Get(k interface{}) (interface{}, bool) {
	stored, found := l.cache.Get(k)
	if !found && l.lookupFn != nil {
		val, ok := l.lookupFn(k)
		if ok {
			l.cache.Add(k, val)
			return val, true
		}

		// not in cache and not return from lookup;
		return nil, false
	}

	return stored, true
}

func (l *LazyLookup) Remove(k interface{}) bool {
	return l.cache.Remove(k)
}
