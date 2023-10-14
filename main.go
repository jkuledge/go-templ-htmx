package main

import (
	"fmt"
	"net/http"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := page(global.Count, 0)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state
	r.ParseForm()

	fmt.Println(r.Form.Get("user"))

	// check if global button has been pressed
	if r.Form.Has("global") {
		global.Count++
	}

	// TODO: update session

	// now that we've udpated state - rerender the page
	getHandler(w, r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postHandler(w, r)
		return
	}
	getHandler(w, r)
}

func main() {
	// component := hello("Josh")

	// http.Handle("/", templ.Handler(component))

	http.HandleFunc("/", rootHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
