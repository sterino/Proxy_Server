// handler/handler.go
package handler

import (
	"app/internal/cache"
	"app/internal/domain/proxy"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	cacheInstance *cache.Cache
}

func NewHandler(cacheInstance *cache.Cache) *Handler {
	return &Handler{cacheInstance: cacheInstance}
}

// @Summary 	request url
// @Tags 	    proxy
// @Description create request
// @Accept 		json
// @Produce 	json
// @Param 		input body  proxy.RequestProxy true "request data"
// @Failure 	400 {object} string
// @Failure 	502 {object} string
// @Router 		/proxy [post]
func (p *Handler) Proxy(c *gin.Context) {
	var request proxy.RequestProxy

	if err := c.BindJSON(&request); err != nil {
		e := p.cacheInstance.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	req, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		e := p.cacheInstance.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e := p.cacheInstance.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	response := p.cacheInstance.Set(uuid.New().String(), request, http.StatusOK, resp.Header)

	c.JSON(http.StatusOK, response)
}

// @Summary get all requests and responses
//
//		@Tags proxy
//		@Description get all history
//		@Accept json
//		@Produce json
//		@Failure 400 {object} string
//	 @Router /caches [get]
func (p *Handler) GetCaches(c *gin.Context) {
	caches, found := p.cacheInstance.GetAll()
	if !found {
		e := p.cacheInstance.SetError(uuid.New().String(), nil, http.StatusNotFound, "Cache not found")
		c.JSON(http.StatusOK, e)
		return
	}
	c.JSON(http.StatusOK, caches)
}

// @Summary get all requests and responses
// @Tags proxy
// @Description get all history
// @Accept json
// @Produce json
// @Param id path string true "request id"
// @Failure 400 {object} string
// @Router /caches/{id} [get]
func (p *Handler) GetCacheById(c *gin.Context) {
	id := c.Param("id")
	resp, req, found := p.cacheInstance.Get(id)
	if !found {
		e := p.cacheInstance.SetError(uuid.New().String(), nil, http.StatusNotFound, "Cache not found")
		c.JSON(http.StatusOK, e)
		return
	}
	c.JSON(http.StatusOK, gin.H{"request": req, "response": resp})
}
