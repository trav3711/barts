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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome!!!"))
	})
	r.Get("/nav", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the navigation page"))
	})

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
