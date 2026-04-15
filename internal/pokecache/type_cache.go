package pokecache

import(
			"time"
			"sync"
			"github.com/ramzygirgis/pokedex/internal/shared_types"
		)


 type cacheEntry struct {
	 createdAt time.Time
	 val shared_types.locationArea // val []byte
 }


type Cache struct {
	cacheMap map[string]cacheEntry
	mux sync.Mutex
	interval time.Duration
}



func NewCache(interval time.Duration) *Cache {
	data := make(map[string]cacheEntry)
	mux := sync.Mutex{}
	c := &Cache{
		cacheMap: data,
		mux: mux,
		interval: interval,
	}
	go c.reapLoop()
	return c
}


func (c *Cache) Add(key string, val shared_types.locationArea) { // val []byte
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cacheMap[key] = cacheEntry{
		val: val,
		createdAt: time.Now(),
	}
}


func (c *Cache) Get(key string) (shared_types.locationArea, bool) { // ([]byte, bool) 
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		b := shared_types.locationArea{} // make([]byte, 0)
		return b, false
	}
	return entry.val, true
}


func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval) 
	defer ticker.Stop()
	for range ticker.C {
		c.mux.Lock()
		now := time.Now()
		for k,v := range c.cacheMap {
			if now.Sub(v.createdAt) >= c.interval {
				delete(c.cacheMap, k)
			}
		}
		c.mux.Unlock()
	}
}
