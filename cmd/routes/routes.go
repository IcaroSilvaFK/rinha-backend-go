package router

import (
	"net/http"

	"github.com/IcaroSilvaFK/rinha-go/cmd/controllers"
	"github.com/IcaroSilvaFK/rinha-go/cmd/services"
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {

	personService := services.NewPersonService()
	personController := controllers.NewPersonController(personService)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/person", personController.Create)
	r.Get("/person", personController.FindPersonBySearchTerm)
	r.Get("/person/{id}", personController.FindPersonById)
	r.Get("/count", personController.CountPersons)
}
