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

// HandlleProxy godoc
// @Summary request url
// @Tags proxy
// @Description create request
// @Accept 	json
// @Produce json
// @Param input body models.RequestProxy true "request data"
// @Success 200 {object} models.ResponseProxy
// @Failure 400 {object} string "Invalid request body"
// @Failure 502 {object} string "Internal server error"
// @Router /proxy [post]
func (p *Proxy) HandleProxy(c *gin.Context) {
	var request models.RequestProxy

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

// GetStore godoc
// @Summary get all requests and responses
// @Tags proxy
// @Description get all history
// @Accept json
// @Produce json
// @Failure 502 {object} string "Internal server error"
// @Router /proxy [get]
func (p *Proxy) GetStore(c *gin.Context) {
	resp, found := p.store.GetAll()
	if !found {
		e := p.store.SetError(uuid.New().String(), nil, http.StatusNotFound, "Data not found")
		c.JSON(http.StatusOK, e)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetStoreById godoc
// @Summary get all requests and responses
// @Tags proxy
// @Description get all history
// @Accept json
// @Produce json
// @Param id path string true "request id"
// @Failure 400 {object} string "Invalid request body"
// @Failure 502 {object} string "Internal server error"
// @Router /proxy/{id} [get]
func (p *Proxy) GetStoreById(c *gin.Context) {
	id := c.Param("id")
	resp, req, found := p.store.Get(id)
	if !found {
		e := p.store.SetError(uuid.New().String(), nil, http.StatusNotFound, "Data not found")
		c.JSON(http.StatusOK, e)
		return
	}
	c.JSON(http.StatusOK, gin.H{"request": req, "response": resp})
}
