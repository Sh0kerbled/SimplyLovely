package main

import (
	"SimplyLovely/internal/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.Register()

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Сервер запущен на localhost:8000")

	http.ListenAndServe(":8000", nil)
}
