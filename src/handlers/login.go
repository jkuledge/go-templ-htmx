package handlers

import (
	"go-templ/src/pages"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		HandlePostLogin(w, r)
		return
	}
	HandleGetLogin(w, r)
}

func HandleGetLogin(w http.ResponseWriter, r *http.Request) {
	pages.Login().Render(r.Context(), w)
}

func HandlePostLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	println("HERE")
}
