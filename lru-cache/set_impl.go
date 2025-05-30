package lru_cache

import (
	"container/list"
	"time"
)

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
