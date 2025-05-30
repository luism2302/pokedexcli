package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMap map[string]cacheEntry
	mut      sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Caching interface {
	Add(key string, value []byte)
	Get(key string) (val []byte, found bool)
	reapLoop()
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		CacheMap: map[string]cacheEntry{},
		mut:      sync.Mutex{},
		interval: interval,
	}
	go newCache.reapLoop()
	return &newCache
}

func (c *Cache) Add(key string, value []byte) {
	newCacheEntry := cacheEntry{CreatedAt: time.Now(), Val: value}
	c.mut.Lock()
	c.CacheMap[key] = newCacheEntry
	c.mut.Unlock()
}

func (c *Cache) Get(key string) (value []byte, found bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	if value, ok := c.CacheMap[key]; ok {
		return value.Val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for passed := range ticker.C {
		c.mut.Lock()
		for url, value := range c.CacheMap {
			if value.CreatedAt.Add(c.interval).Before(passed) {
				delete(c.CacheMap, url)
			} else {
				continue
			}
		}
		c.mut.Unlock()
	}
}
