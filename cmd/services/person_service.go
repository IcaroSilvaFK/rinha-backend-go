package services

import "github.com/IcaroSilvaFK/rinha-go/cmd/models"

type PersonServiceStruct struct {
}

type PersonServiceInterface interface {
	Create(nome, apelido, nascimento, stack string) (string, error)
	FindPersonById(id string) (*models.PersonModel, error)
	FindPersonBySearchTerm(term string) (*[]models.PersonModel, error)
	CountPersons() (int, error)
}

func NewPersonService() PersonServiceInterface {
	return &PersonServiceStruct{}
}

func (*PersonServiceStruct) Create(nome, apelido, nascimento, stack string) (string, error) {

	personId, err := models.CreatePerson(nome, apelido, nascimento, stack)

	if err != nil {
		return "", err
	}

	return personId, nil
}

func (*PersonServiceStruct) FindPersonById(id string) (*models.PersonModel, error) {
	p, err := models.FindPersonById(id)

	return p, err
}

func (*PersonServiceStruct) FindPersonBySearchTerm(term string) (*[]models.PersonModel, error) {
	p, err := models.FindBySearchTerm(term)

	return p, err
}

func (*PersonServiceStruct) CountPersons() (int, error) {

	r, err := models.CountPersons()

	return r, err

}
