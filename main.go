package main

import (
	"embed"
	"fmt"
	"go-templ/src/components"
	"go-templ/src/handlers"
	"go-templ/src/pages"
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
	component := pages.Page(global.Count, userCount)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state
	r.ParseForm()

	// check if global button has been pressed
	if r.Form.Has(components.CountsForm.Global) {
		global.Count++
	}

	if r.Form.Has(components.CountsForm.Session) {
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

var (

	//go:embed dist/output.css
	css embed.FS

	//go:embed htmx/htmx.min.js
	htmx embed.FS
)

func main() {

	// Initialize the session.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := http.NewServeMux()

	mux.Handle("/dist/output.css", http.FileServer(http.FS(css)))
	mux.Handle("/htmx/htmx.min.js", http.FileServer(http.FS(htmx)))

	// Handle POST and GET requests.
	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			time.Sleep(2 * time.Second) // simulate loading
			components.Clicked().Render(r.Context(), w)
		}
	})

	mux.HandleFunc("/login", handlers.HandleLogin)

	// Add the middleware.
	muxWithSessionMiddleware := sessionManager.LoadAndSave(mux)

	// Start the server.
	fmt.Println("listening on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", muxWithSessionMiddleware); err != nil {
		log.Printf("error listening: %v", err)
	}
}
