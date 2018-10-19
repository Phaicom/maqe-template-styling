package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "./templates/home.html", nil)
}
