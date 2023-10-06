package main

import (
	"log"
	"net/http"

	router "github.com/IcaroSilvaFK/rinha-go/cmd/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	router.InitializeRoutes(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
