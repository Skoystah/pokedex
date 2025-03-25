package pokecache

import (
	//"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	//each interval, run reapLoop() to clean the cache
	newCache := Cache{entries: map[string]cacheEntry{}, mu: &sync.Mutex{}}
	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) Add(key string, val []byte) {

	c.mu.Lock()
	defer c.mu.Unlock()

	//fmt.Printf("locked to add : %v", key)
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	return entry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	d := time.Tick(interval)
	for next := range d {
		//remove old entries
		//fmt.Printf("debug - trying to lock and reap - time %v\n", next)
		c.mu.Lock()
		//fmt.Println("locked reaping")
		for key, entry := range c.entries {
			if next.Sub(entry.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
		//fmt.Println("unlocked after reaping")
	}
}
