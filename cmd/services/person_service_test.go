package services_test

import (
	"errors"
	"testing"

	"github.com/IcaroSilvaFK/rinha-go/cmd/database"
	"github.com/IcaroSilvaFK/rinha-go/cmd/models"
	"github.com/IcaroSilvaFK/rinha-go/cmd/services"
	"github.com/go-faker/faker/v4"
)

type user struct {
	Nome    string `json:"nome"`
	Apelido string `json:"apelido"`
	Stack   string `json:"stack"`
}

func TestCreatePerson(t *testing.T) {

	db := database.NewDatabaseConnection()
	m := models.NewPersonModel(db)
	svc := services.NewPersonService(m)

	u := user{}

	faker.FakeData(&u)

	id, err := svc.Create(u.Nome, u.Apelido, "test", u.Stack)

	if !errors.Is(err, nil) {
		t.Errorf("Error on create user %s", err.Error())
	}

	if id == "" {
		t.Error("Invalid uuid")
	}
}

func TestFindUserById(t *testing.T) {

	db := database.NewDatabaseConnection()
	m := models.NewPersonModel(db)
	svc := services.NewPersonService(m)

	u := user{}

	faker.FakeData(&u)

	id, _ := svc.Create(u.Nome, u.Apelido, "test", u.Stack)

	r, err := svc.FindPersonById(id)

	if !errors.Is(err, nil) {
		t.Errorf("Error on find user %s", err.Error())
	}

	if r == nil {
		t.Error("User not found")
	}

}

func TestSearchUserByTerm(t *testing.T) {

	db := database.NewDatabaseConnection()
	m := models.NewPersonModel(db)
	svc := services.NewPersonService(m)

	r, err := svc.FindPersonBySearchTerm("test")

	if !errors.Is(err, nil) {
		t.Errorf("Error on find user %s", err.Error())
	}

	if r == nil {
		t.Error("Users not found")
	}

}

func TestCountUsers(t *testing.T) {
	db := database.NewDatabaseConnection()
	m := models.NewPersonModel(db)
	svc := services.NewPersonService(m)

	r, err := svc.CountPersons()

	if !errors.Is(err, nil) {

		t.Errorf("Error on find user %s", err.Error())
	}

	if r <= 0 {
		t.Error("Count error")
	}

}
