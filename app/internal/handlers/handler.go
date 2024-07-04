package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	validate *validator.Validate

}

func newHandler() *Handler {
	return &Handler{
		validate: validator.New(),

	}
}

func Init(h *Handler) http.Handler {
	router = gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api = router.Group("/api")


	proxy = api.Group("/proxy")
	{
		proxy.POST("", h.)
	}

	return router
}
