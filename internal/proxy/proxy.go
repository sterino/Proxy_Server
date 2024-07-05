package proxy

import (
	"app/internal/cache"
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var cacheInstance *cache.Cache

func init() {
	cacheInstance = cache.NewCache()
}

func Proxy(c *gin.Context) {
	var request model.RequestProxy

	if err := c.BindJSON(&request); err != nil {
		e := cacheInstance.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		e := cacheInstance.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e := cacheInstance.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = v[0]
	}

	response := model.ResponseProxy{
		ID:      uuid.New().String(),
		Status:  resp.StatusCode,
		Headers: header,
		Length:  len(resp.Header),
	}

	cacheInstance.Set(response.ID, response, request)

	c.JSON(http.StatusOK, response)

}

func GetCaches(c *gin.Context) {
	caches := cacheInstance.GetAll()
	c.JSON(http.StatusOK, caches)
}

func GetCacheById(c *gin.Context) {
	id := c.Param("id")
	resp, req, found := cacheInstance.Get(id)
	err := "Cache not found"
	if !found {
		e := cacheInstance.SetError(uuid.New().String(), nil, http.StatusNotFound, err)
		c.JSON(http.StatusOK, e)
		return
	}
	c.JSON(http.StatusOK, gin.H{"request": req, "response": resp})
}
