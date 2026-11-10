package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("web/templates/layouts/base.html")

	tmpl.Execute(w, nil)
}
