package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/aymane-smi/api-test/models"
	"github.com/aymane-smi/api-test/utils"
)

func TestGetAuthorById(t *testing.T){

	utils.InitLogger()

	req, _ := http.Get("http://localhost:8000/author/4")

	if req.StatusCode != 200{
		t.Errorf("invalid request!")
	}
}

func TestAddAuthor(t *testing.T){

	utils.InitLogger()

	author := models.Author{
		Full_name: "test author",
	}

	json, _ := json.Marshal(author)

	req, _ := http.Post("http://localhost:8000/author", "application/json", bytes.NewBuffer(json))

	if req.StatusCode != 200{
		t.Errorf("invalid request!")
	}
}

func TestUpdateAuthor(t *testing.T){
	utils.InitLogger()

	author := models.Author{
		Id: 8,
		Full_name: "test author*",
	}

	json, _ := json.Marshal(author)

	req, _ := http.NewRequest(http.MethodPut, "http://localhost:8000/author", bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "applicatio/json")

	client := &http.Client{}

	resp, _ := client.Do(req)

	if resp.StatusCode != 200{
		t.Errorf("invalid request!")
	}
}

func TestDeleteAuthor(t *testing.T){
	utils.InitLogger()

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8000/author/8", nil)

	client := &http.Client{}

	resp, _ := client.Do(req)

	if resp.StatusCode != 200{
		t.Errorf("invalid request!")
	}



}