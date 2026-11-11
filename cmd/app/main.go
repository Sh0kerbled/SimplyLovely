package main

import (
	"SimplyLovely/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.Register()

	fmt.Println("Сервер запущен на localhost:8000")

	http.ListenAndServe(":8000", nil)
}
