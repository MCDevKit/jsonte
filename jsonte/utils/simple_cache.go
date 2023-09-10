package utils

var cache = map[string]CacheBucket{}

// CacheBucket is a bucket of cache entries
type CacheBucket struct {
	Items map[string]*interface{}
}

// CreateCacheBucket creates a cache bucket
func CreateCacheBucket(bucket string) {
	if _, ok := cache[bucket]; !ok {
		cache[bucket] = CacheBucket{
			Items: map[string]*interface{}{},
		}
	}
}

// PutCache puts a cache entry
func PutCache(bucket string, key string, value interface{}) {
	if b, ok := cache[bucket]; ok {
		b.Items[key] = &value
	} else {
		BadDeveloperError("Cache bucket not found!")
	}
}

// EvictCache evicts a cache entry
func EvictCache(bucket string, key string) {
	if b, ok := cache[bucket]; ok {
		if _, ok := b.Items[key]; ok {
			delete(b.Items, key)
		}
	}
}

// GetCache gets a cache entry
func GetCache(bucket string, key string) *interface{} {
	if b, ok := cache[bucket]; ok {
		if i, ok := b.Items[key]; ok {
			return i
		}
	}
	return nil
}
