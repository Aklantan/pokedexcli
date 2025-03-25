package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu      sync.Mutex
	entries map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		entries: map[string]CacheEntry{},
	}
	go cache.reapLoop(interval)
	return &cache

}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.entries[key]
	if exists {
		return value.val, exists
	} else {
		return nil, exists
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		stale := []string{}
		c.mu.Lock()
		for i, entry := range c.entries {

			if time.Since(entry.createdAt) > interval {
				stale = append(stale, i)

			}

		}

		for _, value := range stale {
			delete(c.entries, value)
		}
		c.mu.Unlock()

	}
}
