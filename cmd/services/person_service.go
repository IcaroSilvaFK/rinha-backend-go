package services

import (
	"errors"

	"github.com/IcaroSilvaFK/rinha-go/cmd/models"
)

type PersonServiceStruct struct {
	model models.PersonModelInterface
}

type PersonServiceInterface interface {
	Create(nome, apelido, nascimento, stack string) (string, error)
	FindPersonById(id string) (*models.PersonModel, error)
	FindPersonBySearchTerm(term string) (*[]models.PersonModel, error)
	CountPersons() (int, error)
}

func NewPersonService(
	model models.PersonModelInterface,
) PersonServiceInterface {
	return &PersonServiceStruct{
		model,
	}
}

func (ps *PersonServiceStruct) Create(nome, apelido, nascimento, stack string) (string, error) {

	personId, err := ps.model.CreatePerson(nome, apelido, nascimento, stack)

	if !errors.Is(err, nil) {
		return "", err
	}

	return personId, nil
}

func (ps *PersonServiceStruct) FindPersonById(id string) (*models.PersonModel, error) {
	p, err := ps.model.FindPersonById(id)

	return p, err
}

func (ps *PersonServiceStruct) FindPersonBySearchTerm(term string) (*[]models.PersonModel, error) {
	p, err := ps.model.FindBySearchTerm(term)

	return p, err
}

func (ps *PersonServiceStruct) CountPersons() (int, error) {

	r, err := ps.model.CountPersons()

	return r, err

}
