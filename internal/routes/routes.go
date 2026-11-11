package routes

import (
	"SimplyLovely/internal/handlers"
	"net/http"
)

func Register() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/store/", handlers.Store)
}
