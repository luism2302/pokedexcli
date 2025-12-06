package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cached   map[string]cacheEntry
	mux      *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Cached:   make(map[string]cacheEntry),
		mux:      &sync.RWMutex{},
		interval: interval,
	}
	go cache.reapLoop()

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mux.Lock()
	cache.Cached[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.mux.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.RLock()
	defer cache.mux.RUnlock()
	if _, ok := cache.Cached[key]; !ok {
		return nil, false
	}
	value := cache.Cached[key].val
	return value, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()
	for range ticker.C {
		cache.mux.Lock()
		for key, entry := range cache.Cached {
			if time.Since(entry.createdAt) > cache.interval {
				delete(cache.Cached, key)
			}
		}
		cache.mux.Unlock()
	}
}
