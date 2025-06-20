package lru_cache

import (
	"container/list"
	"time"
)

// cleanExpired checks expired items and delete them
func (c *LRU) cleanExpired() {

	c.Items.Range(func(k, v interface{}) bool {
		element := v.(*list.Element)
		item := element.Value.(*Item)

		if time.Now().After(item.ExpiresAt) {
			c.Queue.Remove(element)
			c.Items.Delete(k)
		}
		return true
	})
}

// startCleaner starts new ticker and trigger cleanExpired when ticker ticks
func (c *LRU) startCleaner(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanExpired()
		case <-c.stop:
			return
		}
	}
}

func (c *LRU) Close() {
	close(c.stop)
}
