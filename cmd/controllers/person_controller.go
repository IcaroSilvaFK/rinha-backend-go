package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/IcaroSilvaFK/rinha-go/cmd/models"
	"github.com/IcaroSilvaFK/rinha-go/cmd/services"
	"github.com/go-chi/chi/v5"
)

type personController struct {
	svc services.PersonServiceInterface
}

type PersonControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindPersonById(w http.ResponseWriter, r *http.Request)
	FindPersonBySearchTerm(w http.ResponseWriter, r *http.Request)
	CountPersons(w http.ResponseWriter, r *http.Request)
}

func NewPersonController(
	service services.PersonServiceInterface,
) PersonControllerInterface {

	return &personController{
		svc: service,
	}
}

func (ct *personController) Create(w http.ResponseWriter, r *http.Request) {

	var p models.PersonModel

	if err := json.NewDecoder(r.Body).Decode(&p); !errors.Is(err, nil) {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	personId, err := ct.svc.Create(p.Nome, p.Apelido, p.Nascimento, p.Stack)

	if !errors.Is(err, nil) {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("person/%s", personId)))
}

func (ct *personController) FindPersonById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := ct.svc.FindPersonById(id)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (ct *personController) FindPersonBySearchTerm(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")

	if term == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := ct.svc.FindPersonBySearchTerm(term)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func (ct *personController) CountPersons(w http.ResponseWriter, r *http.Request) {

	v, err := ct.svc.CountPersons()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", v)))

}
