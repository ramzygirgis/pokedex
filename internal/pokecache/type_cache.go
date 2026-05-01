package pokecache

import(
			"time"
			"sync"
		)


 type cacheEntry[T any] struct {
	 createdAt time.Time
	 val T
 }


type Cache[T any] struct {
	cacheMap map[string]cacheEntry[T]
	mux sync.Mutex
	interval time.Duration
}



func NewCache[T any](interval time.Duration) *Cache[T] {
	c := &Cache[T]{
		cacheMap: make(map[string]cacheEntry[T]),
		interval: interval,
	}
	go c.reapLoop()
	return c
}


func (c *Cache[T]) Add(key string, val T) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cacheMap[key] = cacheEntry[T]{
		val: val,
		createdAt: time.Now(),
	}
}


func (c *Cache[T]) Get(key string) (T, bool) { // ([]byte, bool) 
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		var zero T
		return zero, false
	}
	return entry.val, true
}


func (c *Cache[T]) reapLoop() {
	ticker := time.NewTicker(c.interval) 
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		now := time.Now()
		for k,v := range c.cacheMap {
			if now.After(v.createdAt.Add(c.interval)) {
				delete(c.cacheMap, k)
			}
		}
		c.mux.Unlock()
	}
}
