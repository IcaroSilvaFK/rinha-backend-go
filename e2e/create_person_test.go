package e2e_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
)

const (
	BASE_URL = "http://localhost:8080"
)

type user struct {
	Nome       string `json:"nome"`
	Apelido    string `json:"apelido"`
	Stack      string `json:"stack"`
	Nascimento string `json:"nascimento"`
}

func TestCreatePerson(t *testing.T) {

	u := user{}

	faker.FakeData(&u)

	body := map[string]string{
		"nome":       u.Nome,
		"apelido":    u.Apelido,
		"nascimento": "teste",
		"stack":      u.Stack,
	}

	bt, _ := json.Marshal(body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/person", BASE_URL), bytes.NewBuffer(bt))

	if !errors.Is(err, nil) {
		t.Errorf("Error on create request %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	r, err := client.Do(req)

	if !errors.Is(err, nil) {
		t.Errorf("Error on create user %s", err.Error())
	}

	if r.StatusCode != 201 {
		t.Error("Invalid status code")
	}

}

func TestFindPersonById(t *testing.T) {

	u := user{}

	faker.FakeData(&u)

	body := map[string]string{
		"nome":       u.Nome,
		"apelido":    u.Apelido,
		"nascimento": "teste",
		"stack":      u.Stack,
	}

	bt, _ := json.Marshal(body)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/person", BASE_URL), bytes.NewBuffer(bt))

	if !errors.Is(err, nil) {
		t.Errorf("Error on create request %s", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	r, _ := client.Do(req)

	response, _ := io.ReadAll(r.Body)

	res, _ := http.Get(fmt.Sprintf("%s/%s", BASE_URL, string(response)))

	if res.StatusCode != 200 {
		t.Error("Invalid status code")
	}

}

func TestCountPerson(t *testing.T) {

	response, err := http.Get(fmt.Sprintf("%s/count", BASE_URL))

	if !errors.Is(err, nil) {
		t.Errorf("Error on create request %s", err.Error())
	}

	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}

	bd, _ := io.ReadAll(response.Body)

	ct := string(bd)

	if ct == "0" {
		t.Error("Invalid count")
	}

}

func TestSearchPersonUsingTerm(t *testing.T) {

	response, err := http.Get(fmt.Sprintf("%s/person?term=test", BASE_URL))

	if !errors.Is(err, nil) {
		t.Errorf("Error on create request %s", err.Error())
	}

	if response.StatusCode != 200 {
		t.Error("Invalid status code")
	}

	bd, _ := io.ReadAll(response.Body)

	var jsonBd any

	err = json.Unmarshal(bd, &jsonBd)

	if !errors.Is(err, nil) {
		t.Errorf("Error on read request %s", err.Error())
	}

	if jsonBd == nil {
		t.Error("Invalid json")
	}
}
