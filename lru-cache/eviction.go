package lru_cache

// evict deletes back element from Queue
func (c *LRU) evict() {
	if element := c.Queue.Back(); element != nil {
		item := c.Queue.Remove(element).(*Item)
		c.Items.Delete(item.Key)
	}
}
