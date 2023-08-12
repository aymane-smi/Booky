package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aymane-smi/api-test/models"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
)

func AddAuthor(w http.ResponseWriter, r *http.Request){
	var author models.Author

	json.NewDecoder(r.Body).Decode(&author)

	msg, err := models.AddAuthor(author)

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

func GetAuthorById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id,_ := strconv.Atoi(vars["id"])

	author := models.GetAuthorById(id)

	if author == nil {
		response := map[string]interface{}{
			"message": "invalid author id",
		}
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusNotFound, jsonResponse)
		return
	}

	response := map[string]interface{}{
		"author": author,
	}
	jsonResponse, _ := json.Marshal(response)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request){
	var author models.Author

	json.NewDecoder(r.Body).Decode(&author)

	a, err := models.UpdateAuthor(author)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"author": a,
	}
	jsonResponse, _ := json.Marshal(response)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	id,_ := strconv.Atoi(vars["id"])

	msg, err := models.DeleteAuthorById(id)

	if err != nil{
		response := map[string]interface{}{
			"message": "invalid author id",
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