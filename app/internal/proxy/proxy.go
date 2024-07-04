package proxy

import (
	"app/internal/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func Proxy(c *gin.Context) {
	var request model.RequestProxy

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Get(request.Url)

	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := time.Now().Format(time.RFC3339Nano)
	headers := make(map[string]string)
	for k, v := range resp.Header {
		headers[k] = v[0]
	}

	response := model.ResponseProxy{
		ID:      id,
		Status:  resp.StatusCode,
		Headers: headers,
		Length:  len(body),
	}

	c.JSON(http.StatusOK, response)

}
