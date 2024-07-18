package router

import (
	"app/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "app/docs"
)

type Router struct {
	proxy *handler.Proxy
}

func NewRouter(proxy *handler.Proxy) *Router {
	return &Router{proxy: proxy}
}

func (r *Router) InitRouters() *gin.Engine {
	router := gin.Default()

	router.POST("/proxy", r.proxy.HandleProxy)
	router.GET("/proxy", r.proxy.GetCaches)
	router.GET("/proxy/:id", r.proxy.GetCacheById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
