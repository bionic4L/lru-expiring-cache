package lru_cache

import (
	"container/list"
	"sync"
	"time"
)

type Item struct {
	Key       string
	Value     interface{}
	ExpiresAt time.Time
}

type LRU struct {
	capacity int
	Queue    *list.List
	Items    sync.Map
	ttl      time.Duration
	stop     chan struct{}
}

// NewLRU creates new LRU cache
func NewLRU(capacity int, ttl, interval time.Duration) *LRU {
	cache := &LRU{
		capacity: capacity,
		Queue:    list.New(),
		Items:    sync.Map{},
		ttl:      ttl,
		stop:     make(chan struct{}),
	}

	go cache.startCleaner(interval)

	return cache
}
