package golru

// Cache is the interface that defines operations
// that a cache can perform.
type Cache interface {
	// Add sets a key and value in the cache. The returned
	// boolean is true if an eviction occured, and false if not.
	Add(interface{}, interface{}) bool

	// Get retrieves the value of a key from the cache. The returned
	// boolean is true if the key was found, and false if not.
	Get(interface{}) (interface{}, bool)

	// Remove deletes a key from the cache. The returned boolean is
	// true if a key was deleted, and false if not.
	Remove(interface{}) bool
}

// New returns a new threadsafe LRU.
func New(size int) Cache {
	return newMultiThreaded(size)
}