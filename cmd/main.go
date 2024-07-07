package main

import (
	"app/internal/cache"
	"app/internal/handler"
	"app/internal/router"
	"fmt"
	"net/http"
)

// @title Proxy Server
// @version 1.0
// @description API Server for Proxy Server App

// @host localhost:8081
// @BasePath /

func main() {

	cacheInstance := cache.NewCache()
	handlerInstance := handler.NewHandler(cacheInstance)
	routerInstance := router.NewRouter(handlerInstance)

	r := routerInstance.InitRouters()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
