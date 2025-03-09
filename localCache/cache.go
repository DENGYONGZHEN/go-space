package localcache

import (
	"sync"
	"time"
)

type Item struct {
	Value      any
	Expiration int64
}

type Cache struct {
	sync.RWMutex
	items      map[string]Item
	gcInterval time.Duration
}

func NewCache(gcInterval time.Duration) *Cache {
	c := &Cache{
		items:      make(map[string]Item),
		gcInterval: gcInterval,
	}
	go c.startGC()
	return c
}

func (c *Cache) startGC() {
	ticker := time.NewTicker(c.gcInterval)
	for {
		<-ticker.C
		c.Lock()
		now := time.Now().UnixNano()
		for k, v := range c.items {
			if now > v.Expiration {
				delete(c.items, k)
			}
		}
		c.Unlock()

	}
}

func (c *Cache) Set(key string, value any, expire time.Duration) {
	c.Lock()
	defer c.Unlock()
	expiration := time.Now().Add(expire).UnixNano()
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.RLock()
	defer c.RUnlock()

	item, exist := c.items[key]
	if !exist {
		return nil, false
	}

	if item.Expiration < time.Now().UnixNano() {
		delete(c.items, key)
		return nil, false
	}
	return item.Value, true
}
