package cache

import (
	"app/internal/model"
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
	c.store.Store(id+"resp", resp)
	c.store.Store(id+"req", req)

	time.AfterFunc(5*time.Second, func() {
		c.store.Delete(id)
	})
}

func (c *Cache) SetError(id string, req interface{}, statusCode int, err string) {
	errMap := make(map[string]string)
	errMap["error"] = err

	resp := model.ResponseProxy{
		ID:      id,
		Status:  statusCode,
		Headers: errMap,
		Length:  len(err),
	}

	c.store.Store(id+"resp", resp)
	c.store.Store(id+"req", req)
}

func (c *Cache) Get(id string) (interface{}, interface{}, bool) {
	res, found := c.store.Load(id + "_resp")
	req, _ := c.store.Load(id + "_req")
	return res, req, found
}
