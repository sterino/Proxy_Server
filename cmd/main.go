package main

import (
	"app/internal/router"
	"fmt"
	"net/http"
)

func main() {

	r := router.Router()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
