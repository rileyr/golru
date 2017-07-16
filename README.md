# golru

golang least recently used cache

---

### Basic Usage


```golang
package main

import(
  "fmt"
  "github.com/rileyr/golru"
)

func main() {
  cache := golru.New(golru.WithSize(30))

  cache.Add("key", "value")
  val, ok := cache.Get("key")
  str := val.(string)

  fmt.Printf("got from cache: %s\n", str)
}
```

### Lazy Lookup

A lookup function can be provided to the initializer. If provided, any GET misses
will attempt to look up the value for the key, store it, and then return it.

```golang
package main

import(
  "fmt"
  "github.com/rileyr/golru"
)

func lookup(k interface{}) (interface{}, bool) {
  // If we want to use this in front of a database, perhaps we are using the ID as the cache key:
  databaseID := k.(int)

  // Get the value from the database (or whatever):
  value := fmt.Sprintf("from-db-%d", databaseID)
  return value, true
}

func main() {
  cache := golru.New(golru.WithSize(30), golru.WithLookup(lookup))
  val, _ := cache.Get(22)
  str := val.(string)
  fmt.Printf("got from cache: %s\n", str)
}
```
