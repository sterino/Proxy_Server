// handler/proxy.go
package handler

import (
	"app/internal/domain/proxy"
	"app/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Proxy struct {
	store *store.Store
}

func NewHandler(store *store.Store) *Proxy {
	return &Proxy{store: store}
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
func (p *Proxy) HandleProxy(c *gin.Context) {
	var request proxy.RequestProxy

	if err := c.BindJSON(&request); err != nil {
		e := p.store.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	req, err := http.NewRequest(request.Method, request.URL, nil)
	if err != nil {
		e := p.store.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		e := p.store.SetError(uuid.New().String(), request, http.StatusBadRequest, err.Error())
		c.JSON(http.StatusOK, e)
		return
	}

	response := p.store.Set(uuid.New().String(), request, http.StatusOK, resp.Header)

	c.JSON(http.StatusOK, response)
}

// @Summary get all requests and responses
//
//		@Tags proxy
//		@Description get all history
//		@Accept json
//		@Produce json
//		@Failure 400 {object} string
//	 @Router /proxy [get]
func (p *Proxy) GetStore(c *gin.Context) {
	caches, found := p.store.GetAll()
	if !found {
		e := p.store.SetError(uuid.New().String(), nil, http.StatusNotFound, "Cache not found")
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
// @Router /proxy/{id} [get]
func (p *Proxy) GetStoreById(c *gin.Context) {
	id := c.Param("id")
	resp, req, found := p.store.Get(id)
	if !found {
		e := p.store.SetError(uuid.New().String(), nil, http.StatusNotFound, "Cache not found")
		c.JSON(http.StatusOK, e)
		return
	}
	c.JSON(http.StatusOK, gin.H{"request": req, "response": resp})
}
