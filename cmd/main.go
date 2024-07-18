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

	stores := store.NewStore()
	handlers := handler.NewHandler(stores)
	routers := router.NewRouter(handlers)

	r := routers.InitRouters()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
