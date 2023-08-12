package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aymane-smi/api-test/models"
	prometheus_book "github.com/aymane-smi/api-test/prometheus"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
)

func GetBookById(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()


	book := models.GetBookById(vars["id"])

	if book == nil {
		response := map[string]interface{}{
			"message": "invalid book id",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("controller", "GetBookById", "controllers").Inc()
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

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	json.NewDecoder(r.Body).Decode(&book)

	msg, err := models.AddBook(book)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("controller", "AddBook", "controllers").Inc()
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

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	json.NewDecoder(r.Body).Decode(&book)

	b, err := models.UpdateBook(book)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("controller", "UpdateBook", "controllers").Inc()
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

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	msg, err := models.DeleteById(vars["id"])

	if err != nil{
		response := map[string]interface{}{
			"message": "invalid book id",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("controller", "DeleteBook", "controllers").Inc()
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