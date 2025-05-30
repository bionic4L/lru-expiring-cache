package lru_cache

func (c *LRU) evict() {
	if element := c.Queue.Back(); element != nil {
		item := c.Queue.Remove(element).(*Item)
		c.Items.Delete(item.Key)
	}
}
