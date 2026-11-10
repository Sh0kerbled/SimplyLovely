package routes

import (
	"SimplyLovely/internal/handlers"
	"net/http"
)

func Register() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/store/", handlers.Store)
}
