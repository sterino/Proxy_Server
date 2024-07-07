package router

import (
	"app/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "app/docs"
)

type Router struct {
	proxy *handler.Handler
}

func NewRouter(proxy *handler.Handler) *Router {
	return &Router{proxy: proxy}
}

func (r *Router) InitRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/proxy", r.proxy.Proxy)
	router.GET("/caches", r.proxy.GetCaches)
	router.GET("/caches/:id", r.proxy.GetCacheById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
