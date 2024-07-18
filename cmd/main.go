package main

import (
	"app/internal/handler"
	"app/internal/router"
	"app/internal/store"
	"fmt"
	"net/http"
)

// @title Proxy Server
// @version 1.0
// @description API Server for Proxy Server App

// @host localhost:8080
// @BasePath /
func main() {

	cacheInstance := store.NewCache()
	handlerInstance := handler.NewHandler(cacheInstance)
	routerInstance := router.NewRouter(handlerInstance)

	r := routerInstance.InitRouters()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
