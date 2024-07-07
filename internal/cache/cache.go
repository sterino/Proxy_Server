// cache/cache.go
package cache

import (
	"app/internal/domain/proxy"
	"net/http"
	"sync"
)

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Set(id string, req proxy.RequestProxy, statusCode int, head http.Header) proxy.ResponseProxy {

	header := make(map[string]string)
	for k, v := range head {
		header[k] = v[0]
	}

	resp := proxy.ResponseProxy{
		ID:      id,
		Status:  statusCode,
		Headers: header,
		Length:  len(head),
	}

	c.store.Store(id+"_resp", resp)
	c.store.Store(id+"_req", req)

	return resp
}

func (c *Cache) SetError(id string, req interface{}, statusCode int, err string) proxy.Error {
	errMap := make(map[string]string)
	errMap["error"] = err

	resp := proxy.Error{
		ID:      id,
		Status:  statusCode,
		Headers: errMap,
		Message: err,
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

func (c *Cache) GetAll() (map[string]interface{}, bool) {
	allCache := make(map[string]interface{})
	c.store.Range(func(k, v interface{}) bool {
		allCache[k.(string)] = v
		return true
	})

	return allCache, len(allCache) > 0
}
