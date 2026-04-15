package pokecache

import(
			"time"
			"sync"
		)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


type Cache struct {
	cache map[string]cacheEntry
	mux sync.Mutex
	interval time.Duration
}



func NewCache(interval time.Duration) *Cache {
	data := make(map[string]cacheEntry)
	mux := sync.Mutex{}
	c := &Cache{
		cache: data,
		mux: mux,
		interval: interval,
	}
	go c.reapLoop()
	return c
}


func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = cacheEntry{
		val: val,
		createdAt: time.Now()
	}
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	if !ok {
		b := make([]byte, 0)
		return b, false
	}
	return entry.val, true
}


func (c *Cache) reapLoop() {
	ticker := time.Ticker(c.inverval) 
	for range ticker.C {
		c.mux.Lock()
		now = time.Now()
		for k,v := range c.cache {
			if now.Sub(v.createdAt) >= c.interval {
				delete(c.cache, k)
			}
		}
		c.mux.Unlock()
	}
}
