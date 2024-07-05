package proxy

import (
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Proxy(c *gin.Context) {
	var request model.RequestProxy

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, response)

}
