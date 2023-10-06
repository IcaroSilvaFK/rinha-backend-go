package models_test

import (
	"errors"
	"testing"

	"github.com/IcaroSilvaFK/rinha-go/cmd/database"
	"github.com/IcaroSilvaFK/rinha-go/cmd/models"
	"github.com/IcaroSilvaFK/rinha-go/cmd/utils"
	"github.com/go-faker/faker/v4"
)

type user struct {
	Nome string 
	Apelido string 
	Nascimento string 
	Stack string
}

func TestCreatePerson(t *testing.T) {
	
	db := database.NewDatabaseConnection()
	p := models.NewPersonModel(db)

	u := user{}

	 faker.FakeData(&u)
	id, err := p.CreatePerson(u.Nome, u.Apelido, "test", u.Stack)

	if !errors.Is(err, nil) {
		t.Errorf("Error on create user %s", err.Error())
	}


	if !utils.IsValidUUID(id) {
		t.Error("Invalid uuid")
	}

}

func TestFindBySearchTerm(t *testing.T) {


	db := database.NewDatabaseConnection()
	p := models.NewPersonModel(db)

	u, err := p.FindBySearchTerm("test")


	if !errors.Is(err, nil) {
		t.Errorf("Error on find user %s", err.Error())
	}


	if u == nil {
		t.Error("User not found")
	}
}


func TestFindUserById(t *testing.T) {
	db := database.NewDatabaseConnection()
	p := models.NewPersonModel(db)

	u := user{}

	faker.FakeData(&u)

	id, _ := p.CreatePerson(u.Nome, u.Apelido, "test", u.Stack)


	r, err := p.FindPersonById(id)


	if !errors.Is(err, nil) {
		t.Errorf("Error on find user %s", err.Error())
	}

	if r == nil {
		t.Error("User not found")
	}
}



func TestCountPersons(t *testing.T) {
	db := database.NewDatabaseConnection()
	p := models.NewPersonModel(db)


	r, err := p.CountPersons()

	if !errors.Is(err, nil) {
		t.Errorf("Error on find user %s", err.Error())
	}

	if r <= 0 {
		t.Error("Count error")
	}
}