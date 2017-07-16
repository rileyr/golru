# golru

golang least recently used cache

---

### Basic Usage


```golang
package main

import(
  "fmt"
  "github.com/rileyr/golru/lru"
)

func main() {
  cache := lru.New(lru.WithSize(30))

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
  "github.com/rileyr/golru/lru"
)

func lookup(k interface{}) (interface{}, bool) {
  // If we want to use this in front of a database, perhaps we are using the ID as the cache key:
  databaseID := k.(int)

  // Get the value from the database (or whatever):
  value := fmt.Sprintf("from-db-%d", databaseID)
  return value, true
}

func main() {
  cache := lru.New(lru.WithSize(30), lru.WithLookup(lookup))
  val, _ := cache.Get(22)
  str := val.(string)
  fmt.Printf("got from cache: %s\n", str)
}
```

### Standalone HTTP LRU

This repo provides a standalone HTTP server that serves the cache as json. To use:

```shell
$> go install github.com/rileyr/golru/weblru
$> weblru -port=:3015 -username=basicauthusername -password=basicauthpassword -size=100
```
