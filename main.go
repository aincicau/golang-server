package main

import (
	"golangserver/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	http.HandleFunc("/v1/", rest.Welcome)
	http.HandleFunc("/v1/bye", rest.Welcome)
	chi := chi.NewRouter()
	http.ListenAndServe(":8080", nil)
}
