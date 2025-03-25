package pokecache

import (
	"slices"
	"testing"
	"time"
)

func TestCacheGet(t *testing.T) {
	cache := NewCache(time.Second * 5)

	cacheEntries := map[string]cacheEntry{
		"url1": cacheEntry{time.Now(), []byte("123")},
	}
	cache.entries = cacheEntries

	//get val
	val, exists := cache.Get("url1")
	if !exists {
		t.Errorf("Get failed - value for url1 does not exist")
	}
	if !slices.Equal(val, cache.entries["url1"].val) {
		t.Errorf("Get failed - value retrieve %v does not equal value expected %v", val, cache.entries["url1"].val)
	}
}
