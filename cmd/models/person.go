package models

import (
	"errors"

	"github.com/IcaroSilvaFK/rinha-go/cmd/database"
	"github.com/IcaroSilvaFK/rinha-go/cmd/utils"
)

type PersonModel struct {
	Nome       string `json:"name"`
	Apelido    string `json:"apelido"`
	Id         string `json:"-"`
	Nascimento string `json:"nascimento"`
	Stack      string `json:"stack"`
}

func CreatePerson(nome, apelido, nascimento, stack string) (string, error) {

	personId := utils.NewUUID()

	p := &PersonModel{
		Nome:       nome,
		Apelido:    apelido,
		Id:         personId,
		Nascimento: nascimento,
		Stack:      stack,
	}

	sqlDB := database.NewDatabaseConnection()

	stmt, err := sqlDB.Prepare("INSERT INTO pessoas (id, nome, apelido, nascimento, stack) VALUES ($1, $2, $3, $4, $5)")

	if !errors.Is(err, nil) {
		return "", err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.Nome, p.Apelido, p.Nascimento, p.Stack)

	if !errors.Is(err, nil) {
		return "", err
	}

	return p.Id, nil
}

func FindPersonById(id string) (*PersonModel, error) {

	sqlDB := database.NewDatabaseConnection()

	stmt, err := sqlDB.Prepare("SELECT apelido,nome,nascimento,stack FROM pessoas WHERE id = $1")

	if !errors.Is(err, nil) {
		return nil, err
	}

	r, err := stmt.Query(id)

	if !errors.Is(err, nil) {
		return nil, err
	}

	defer r.Close()

	var p PersonModel

	if !r.Next() {
		return nil, nil
	}

	err = r.Scan(&p.Apelido, &p.Nome, &p.Nascimento, &p.Stack)
	if !errors.Is(err, nil) {

		return nil, err
	}

	return &p, nil
}

func FindBySearchTerm(term string) (*[]PersonModel, error) {

	sqlDB := database.NewDatabaseConnection()

	stmt, err := sqlDB.Prepare("SELECT apelido,nome,nascimento,stack FROM pessoas WHERE busca ilike '%' || $1 || '%' limit 50")

	if !errors.Is(err, nil) {
		return nil, err
	}

	r, err := stmt.Query(term)

	if !errors.Is(err, nil) {
		return nil, err
	}

	defer r.Close()

	p := []PersonModel{}

	for r.Next() {

		person := new(PersonModel)
		r.Scan(&person.Apelido, &person.Nome, &person.Nascimento, &person.Stack)

		p = append(p, *person)
	}

	return &p, nil
}
