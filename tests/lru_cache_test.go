package lru_cache_test

import (
	lruCache "lru-and-lfu-cache/lru-cache"
	"testing"
	"time"
)

func TestLRUCache(t *testing.T) {
	type testCase struct {
		name       string
		setup      func(*lruCache.LRU)
		key        string
		want       interface{}
		sleepAfter time.Duration
	}

	tests := []testCase{
		{
			name: "Test set and get",
			setup: func(c *lruCache.LRU) {
				c.Set("a", "apple")
			},
			key:  "a",
			want: "apple",
		},
		{
			name: "Test expired item",
			setup: func(c *lruCache.LRU) {
				c.Set("b", "I'm gonna die")
			},
			key:        "b",
			want:       nil,
			sleepAfter: 2 * time.Second,
		},
		{
			name: "Test eviction",
			setup: func(c *lruCache.LRU) {
				c.Set("a", 1) // Должен быть вытеснен
				c.Set("b", 2)
				c.Set("c", 3)
				c.Set("d", 4)
			},
			key:  "a",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := lruCache.NewLRU(3, 1*time.Second, 2*time.Second)
			tt.setup(cache)
			if tt.sleepAfter > 0 {
				time.Sleep(tt.sleepAfter)
			}
			got := cache.Get(tt.key)
			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
