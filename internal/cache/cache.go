package cache

import (
	"sync"
	"time"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(id string, resp interface{}, req interface{}) {
	c.store.Store(id, resp)
	c.store.Store(id, req)
	time.AfterFunc(5*time.Second, func() {
		c.store.Delete(id)
	})
}

func (c *Cache) Get(id string) (interface{}, bool) {
	return c.store.Load(id)
}
