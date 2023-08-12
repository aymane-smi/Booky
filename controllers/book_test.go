package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aymane-smi/api-test/models"
	"github.com/aymane-smi/api-test/utils"
)

func TestGetBookById(t *testing.T){

	utils.InitLogger()

	req, _ := http.Get("http://localhost:8000/book/21a1e439-baf5-400a-b515-c10ed7ead9b5")

	if req.StatusCode != 200{
		t.Errorf("invalid request!")
	}
}

func TestAddBook(t *testing.T){

	utils.InitLogger()

	book := models.Book{
		Title: "test book",
		ISBF: "TEST ISBF",
		Page: 100,
		Author: 4,
	}

	json, _ := json.Marshal(book)

	req, _ := http.Post("http://localhost:8000/book", "application/json", bytes.NewBuffer(json))

	if req.StatusCode != 200{
		t.Errorf("invalid request!")
	}
}

func TestUpdateBook(t *testing.T){
	utils.InitLogger()

	book := models.Book{
		Id: "2d1a8ebd-9d9c-4122-87f6-d5e305d6e46e",
		Title: "test book**",
		ISBF: "TEST ISBF**",
		Page: 102,
		Author: 5,
	}

	json, _ := json.Marshal(book)

	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8000/book", bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "applicatio/json")

	client := &http.Client{}

	resp, _ := client.Do(req)

	if resp.StatusCode != 200{
		t.Errorf("invalid request!")
	}
}

func TestDeleteBook(t *testing.T){
	utils.InitLogger()

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8000/book/2d1a8ebd-9d9c-4122-87f6-d5e305d6e46e", nil)

	client := &http.Client{}

	resp, _ := client.Do(req)

	if resp.StatusCode != 200{
		t.Errorf("invalid request!")
	}

}