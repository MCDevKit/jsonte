package utils

import "sync"

var (
	cache   = map[string]*CacheBucket{}
	cacheMu sync.RWMutex
)

// CacheBucket is a bucket of cache entries
type CacheBucket struct {
	Items map[string]*interface{}
	mu    sync.RWMutex
}

// CreateCacheBucket creates a cache bucket
func CreateCacheBucket(bucket string) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if _, ok := cache[bucket]; !ok {
		cache[bucket] = &CacheBucket{
			Items: map[string]*interface{}{},
		}
	}
}

// PutCache puts a cache entry
func PutCache(bucket, key string, value interface{}) {
	cacheMu.RLock()
	b, ok := cache[bucket]
	cacheMu.RUnlock()
	if ok {
		b.mu.Lock()
		b.Items[key] = &value
		b.mu.Unlock()
	} else {
		BadDeveloperError("Cache bucket not found!")
	}
}

// EvictCache evicts a cache entry
func EvictCache(bucket, key string) {
	cacheMu.RLock()
	b, ok := cache[bucket]
	cacheMu.RUnlock()
	if ok {
		b.mu.Lock()
		if _, ok := b.Items[key]; ok {
			delete(b.Items, key)
		}
		b.mu.Unlock()
	}
}

// GetCache gets a cache entry
func GetCache(bucket, key string) *interface{} {
	cacheMu.RLock()
	b, ok := cache[bucket]
	cacheMu.RUnlock()
	if ok {
		b.mu.RLock()
		i, ok := b.Items[key]
		b.mu.RUnlock()
		if ok {
			return i
		}
	}
	return nil
}
