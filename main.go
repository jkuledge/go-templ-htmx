package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

type GlobalState struct {
	Count int
}

var global GlobalState
var sessionManager *scs.SessionManager

func getHandler(w http.ResponseWriter, r *http.Request) {
	userCount := sessionManager.GetInt(r.Context(), "count")
	component := page(global.Count, userCount)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state
	r.ParseForm()

	// check if global button has been pressed
	if r.Form.Has("global") {
		global.Count++
	}

	if r.Form.Has("user") {
		currentCount := sessionManager.GetInt(r.Context(), "count")
		sessionManager.Put(r.Context(), "count", currentCount+1)
	}

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

	// initialize static dir
	fs := http.FileServer(http.Dir("./htmx"))
	http.Handle("/htmx/", http.StripPrefix("/htmx/", fs))

	// Initialize the session.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := http.NewServeMux()

	// Handle POST and GET requests.
	mux.HandleFunc("/", rootHandler)

	// Add the middleware.
	muxWithSessionMiddleware := sessionManager.LoadAndSave(mux)

	// Start the server.
	fmt.Println("listening on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", muxWithSessionMiddleware); err != nil {
		log.Printf("error listening: %v", err)
	}
}
