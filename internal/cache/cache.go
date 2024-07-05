package cache

import (
	"app/internal/model"
	"sync"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(id string, resp model.ResponseProxy, req model.RequestProxy) {
	c.store.Store(id+"_resp", resp)
	c.store.Store(id+"_req", req)
}

func (c *Cache) SetError(id string, req interface{}, statusCode int, err string) interface{} {
	errMap := make(map[string]string)
	errMap["error"] = err

	resp := model.ResponseProxy{
		ID:      id,
		Status:  statusCode,
		Headers: errMap,
		Length:  len(err),
	}
	c.store.Store(id+"_resp", resp)
	c.store.Store(id+"_req", req)

	return resp
}

func (c *Cache) Get(id string) (interface{}, interface{}, bool) {
	res, found := c.store.Load(id + "_resp")
	req, _ := c.store.Load(id + "_req")

	return res, req, found
}

func (c *Cache) GetAll() map[string]interface{} {
	allCache := make(map[string]interface{})
	c.store.Range(func(k, v interface{}) bool {
		allCache[k.(string)] = v
		return true
	})

	return allCache
}
