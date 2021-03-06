package main

import (
	"golangserver/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	http.HandleFunc("/", rest.Welcome)
	http.HandleFunc("/health", rest.Welcome)
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/person", rest.GetPerson)
		r.Post("/person", rest.PostPerson)
	})
	http.ListenAndServe(":8080", router)
}
