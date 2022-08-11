package utils

import "time"

var cache = map[string]CacheBucket{}

type CacheBucket struct {
	Items           map[string]CacheEntry
	DefaultLifetime int64
}

type CacheEntry struct {
	Value    *interface{}
	ExpireAt int64
}

func CreateCacheBucket(bucket string, defaultLifetime int64) {
	if _, ok := cache[bucket]; !ok {
		cache[bucket] = CacheBucket{
			DefaultLifetime: defaultLifetime,
			Items:           map[string]CacheEntry{},
		}
	}
}

func PutCache(bucket string, key string, value interface{}) {
	if b, ok := cache[bucket]; ok {
		b.Items[key] = CacheEntry{
			Value:    &value,
			ExpireAt: time.Now().Unix() + cache[bucket].DefaultLifetime,
		}
	} else {
		panic("Cache bucket not found!")
	}
}

func EvictCache(bucket string, key string) {
	if b, ok := cache[bucket]; ok {
		if _, ok := b.Items[key]; ok {
			delete(b.Items, key)
		}
	}
}

func GetCache(bucket string, key string) *interface{} {
	if b, ok := cache[bucket]; ok {
		if i, ok := b.Items[key]; ok {
			if i.ExpireAt > time.Now().Unix() {
				return cache[bucket].Items[key].Value
			} else {
				EvictCache(bucket, key)
			}
		}
	}
	return nil
}

func ClearCache(bucket string) {
	delete(cache, bucket)
}

func ClearAllCache() {
	cache = map[string]CacheBucket{}
}
