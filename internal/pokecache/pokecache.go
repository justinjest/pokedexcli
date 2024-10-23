package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var newEntry = &cacheEntry{}
	newEntry.createdAt = time.Now()
	newEntry.val = value
	c.cache[key] = *newEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	value, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		func() {
			for key, value := range c.cache {
				if time.Since(value.createdAt) >= interval {
					delete(c.cache, key)
				}
			}
		}()
		c.mu.Unlock()
	}
}
