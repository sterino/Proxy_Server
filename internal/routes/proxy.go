package routes

import (
	"app/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "app/docs"
)

func InitRouters(proxy *handler.Proxy) *gin.Engine {
	router := gin.Default()

	router.POST("/proxy", proxy.HandleProxy)
	router.GET("/proxy", proxy.GetStore)
	router.GET("/proxy/:id", proxy.GetStoreById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
