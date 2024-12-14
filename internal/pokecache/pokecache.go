package pokecache

import (
	"fmt"
	"sync"
	"time"
)

// CacheEntry struct holds the creation time and the value of the cache entry
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache struct holds map of cache entries and a mutex for safe access
type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        *sync.RWMutex
}

// Add - adds a new cache entry to the cache
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get - retrieves a cache entry from the cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
    fmt.Printf("LOG: Reaping cache entries\n")
    c.mutex.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) >= interval {
				delete(c.cacheEntries, key)
			}
		}
    fmt.Printf("LOG: Reaping complete\n")
    c.mutex.Unlock()
	}
}

// NewCache - creates a new cache with a reaping interval
func NewCache(interval time.Duration) *Cache {
	c := Cache{
		cacheEntries: map[string]cacheEntry{},
		mutex:        &sync.RWMutex{},
	}
	go c.reapLoop(interval)
	return &c
}
