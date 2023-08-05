package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aymane-smi/api-test/models"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
)

func GetBookById(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	book := models.GetBookById(vars["id"])

	if book == nil {
		response := map[string]interface{}{
			"message": "invalid book id",
		}
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusNotFound, jsonResponse)
		return
	}

	response := map[string]interface{}{
		"book": book,
	}
	jsonResponse, _ := json.Marshal(response)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func AddBook(w http.ResponseWriter, r *http.Request){
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	msg, err := models.AddBook(book)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"message": msg,
	}
	jsonResponse, _ := json.Marshal(response)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	b, err := models.UpdateBook(book)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"book": b,
	}
	jsonResponse, _ := json.Marshal(response)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	msg, err := models.DeleteById(vars["id"])

	if err != nil{
		response := map[string]interface{}{
			"message": "invalid book id",
		}
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"message": msg,
	}
	jsonResponse, _ := json.Marshal(response)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}