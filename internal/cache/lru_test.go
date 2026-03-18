package cache

import (
	"strconv"
	"testing"
)

func BenchmarkLRUCachePut(b *testing.B) {
	cache := NewLRUCache(1024)
	value := []byte("benchmark-value")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Put(strconv.Itoa(i), value)
	}
}

func BenchmarkLRUCacheGetHit(b *testing.B) {
	cache := NewLRUCache(1024)
	for i := 0; i < 1024; i++ {
		cache.Put(strconv.Itoa(i), []byte("value"))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = cache.Get(strconv.Itoa(i % 1024))
	}
}

func BenchmarkLRUCacheGetMiss(b *testing.B) {
	cache := NewLRUCache(1024)
	for i := 0; i < 1024; i++ {
		cache.Put(strconv.Itoa(i), []byte("value"))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = cache.Get("missing-key-" + strconv.Itoa(i))
	}
}
