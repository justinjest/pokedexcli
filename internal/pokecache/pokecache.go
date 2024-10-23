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

func (c *Cache) Add(key string, value []byte) {
	var newEntry = &cacheEntry{}
	newEntry.createdAt = time.Now()
	newEntry.val = value
	c.cache[key] = *newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {

}

func (c *Cache) reapLoop() {

}
