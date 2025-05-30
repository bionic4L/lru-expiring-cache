package lru_cache

import (
	"container/list"
	"fmt"
	"github.com/robfig/cron/v3"
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
}

func NewLRU(capacity int, ttl time.Duration, interval string) *LRU {
	cache := &LRU{
		capacity: capacity,
		Queue:    list.New(),
		Items:    sync.Map{},
		ttl:      ttl,
	}

	c := cron.New()
	if _, err := c.AddFunc(interval, cache.cleanExpired); err != nil {
		fmt.Println(err)
	}
	c.Start()

	return cache
}

func (c *LRU) Set(key string, value interface{}) bool {
	if v, exists := c.Items.Load(key); exists {
		element := v.(*list.Element)
		item := element.Value.(*Item)

		item.ExpiresAt.Add(c.ttl)

		item.Value = value
		c.Queue.MoveToFront(element)

		return true
	}

	if c.Queue.Len() == c.capacity {
		c.evict()
	}

	item := &Item{
		Key:       key,
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
	}

	listElement := c.Queue.PushFront(item)
	c.Items.Store(item.Key, listElement)

	return true
}

func (c *LRU) Get(key string) interface{} {
	value, exists := c.Items.Load(key)
	if !exists {
		return nil
	}

	element := value.(*list.Element)
	item := element.Value.(*Item)

	if time.Now().After(item.ExpiresAt) {
		c.Queue.Remove(element)
		c.Items.Delete(key)
	}

	c.Queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c *LRU) evict() {
	if element := c.Queue.Back(); element != nil {
		item := c.Queue.Remove(element).(*Item)
		c.Items.Delete(item.Key)
	}
}

func (c *LRU) cleanExpired() {
	var toDelete []string

	c.Items.Range(func(k, v interface{}) bool {
		element := v.(*list.Element)
		item := element.Value.(*Item)

		if time.Now().After(item.ExpiresAt) {
			c.Queue.Remove(element)
			toDelete = append(toDelete, item.Key)
		}
		return true
	})

	for _, k := range toDelete {
		c.Items.Delete(k)
	}
}
