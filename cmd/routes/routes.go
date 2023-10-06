package router

import (
	"net/http"

	"github.com/IcaroSilvaFK/rinha-go/cmd/controllers"
	"github.com/IcaroSilvaFK/rinha-go/cmd/database"
	"github.com/IcaroSilvaFK/rinha-go/cmd/models"
	"github.com/IcaroSilvaFK/rinha-go/cmd/services"
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {

	db := database.NewDatabaseConnection()
	personModel := models.NewPersonModel(db)
	personService := services.NewPersonService(personModel)
	personController := controllers.NewPersonController(personService)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/person", personController.Create)
	r.Get("/person", personController.FindPersonBySearchTerm)
	r.Get("/person/{id}", personController.FindPersonById)
	r.Get("/count", personController.CountPersons)
}
