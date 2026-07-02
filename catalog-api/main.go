package main

import (
	"context"
	"log"
	"net/http"

	"catalog-api/db"

	"github.com/go-chi/chi/v5"
)

func main() {
	client := db.Connect()
	defer client.Disconnect(context.Background())

	r := chi.NewRouter()

	r.Get("/events", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("events list"))
	})

	log.Println("listening on :8080")
	http.ListenAndServe(":8080", r)
}
