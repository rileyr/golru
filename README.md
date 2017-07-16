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

This repo provides a standalone HTTP server that serves the cache as json. Flags that can be
provided on startup:
  - `size`: Upper limit on number of cached items. (Default 100)
  - `port`: Port to serve the cache from. (Default :3030)
  - `username` : BasicAuth username.
  - `password` : BasicAuth password.

If a username *and* password are not provided, no authentication check will occur.

### Usage:
```shell
$> go install github.com/rileyr/golru/weblru
$> weblru -port=:3015  -size=100
```

Requests can be made into the cache:
```
POST localhost:3015/cache:
{
  "key": "key",
  "value": "some value"
}
#=> 201 CREATED
#=> 400 UNPROCESSABLE ENTITY
#=> 500 INTERNAL SERVER ERROR


GET localhost:3015/cache?key=foo:
{
  "key": "key",
  "value": "some value"
}
#=> 200 FOUND
#=> 400 UNPROCESSABLE ENTITY
#=> 500 INTERNAL SERVER ERROR

DELETE localhost:3015/cache:
{
  "key": "key",
}

#=> 200 OK
```
