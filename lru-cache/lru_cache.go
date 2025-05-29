package lru_cache

import (
	"container/list"
	"sync"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	Queue    *list.List
	Items    sync.Map
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		Queue:    list.New(),
		Items:    sync.Map{},
	}
}

func (c *LRU) Set(key string, value interface{}) bool {
	if v, exists := c.Items.Load(key); exists {
		element := v.(*list.Element)
		item := element.Value.(*Item)

		item.Value = value
		c.Queue.PushFront(element)

		return true
	}

	if c.Queue.Len() == c.capacity {
		c.evict()
	}

	item := &Item{
		Key:   key,
		Value: value,
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
	c.Queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c *LRU) evict() {
	if element := c.Queue.Back(); element != nil {
		item := c.Queue.Remove(element).(*Item)
		c.Items.Delete(item.Key)
	}
}
