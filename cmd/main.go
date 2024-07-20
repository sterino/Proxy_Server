package main

import (
	"app/internal/handler"
	"app/internal/routes"
	"app/internal/store"
	"net/http"
)

// @title Proxy Server
// @version 1.0
// @description API Server for Proxy Server App

// @BasePath /
func main() {

	stores := store.NewStore()
	handlers := handler.NewHandler(stores)

	r := routes.InitRouters(handlers)

	http.ListenAndServe(":8080", r)
}
