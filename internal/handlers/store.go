package handlers

import (
	"fmt"
	"net/http"
)

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "тут будет весь ассортимент магазина")
}
