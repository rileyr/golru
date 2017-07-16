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
    cache := golru.New(5)
    
    for i:=0;i<=30;i++ {
        addToCache(cache, i)
    }
}

func addToCache(c golru.Cache, n int) {
    key := fmt.Sprintf("key-%s", n)
    evicted := c.Add(key, n)
    fmt.Printf("added: %n, eviction occurred: %s", key, evicted)
}
```

