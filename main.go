package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// main is the function executed when the binary is run.
func main() {

	// Default port to 8080, but also allows override from env vars.
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	// Router is what decides what paths do what.
	r := chi.NewRouter()

	// GET / -> function
	r.Get("/", routeHandler)

	r.Get("/nav", listPosts)

	r.Get("/post/{postId}", getPost)

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func routeHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome!!!"))
}

func listPosts(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is the navigation page"))
}

func getPost(w http.ResponseWriter, req *http.Request) {
	post := chi.URLParam(req, "postId")

	w.Write([]byte(post))
}
