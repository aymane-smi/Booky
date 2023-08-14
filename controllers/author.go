package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/aymane-smi/api-test/models"
	prometheus_book "github.com/aymane-smi/api-test/prometheus"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
)

func AddAuthor(w http.ResponseWriter, r *http.Request){
	//timestamp of the handler being called
	start := time.Now()
	var author models.Author

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	json.NewDecoder(r.Body).Decode(&author)

	msg, err := models.AddAuthor(author)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("author", "AddAuthor", "controllers").Inc()
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"message": msg,
	}
	jsonResponse, _ := json.Marshal(response)
	//get the remaining timefrom the start and give it to the histogram
	duration := time.Since(start).Seconds()
	prometheus_book.RequestDuration.Observe(duration)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request){
	//timestamp of the handler being called
	start := time.Now()
	vars := mux.Vars(r)

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	id,_ := strconv.Atoi(vars["id"])

	author := models.GetAuthorById(id)

	if author == nil {
		response := map[string]interface{}{
			"message": "invalid author id",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("author", "GetAuthorById", "controllers").Inc()
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusNotFound, jsonResponse)
		return
	}

	response := map[string]interface{}{
		"author": author,
	}
	jsonResponse, _ := json.Marshal(response)
	//get the remaining timefrom the start and give it to the histogram
	duration := time.Since(start).Seconds()
	prometheus_book.RequestDuration.Observe(duration)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request){
	//timestamp of the handler being called
	start := time.Now()
	var author models.Author

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	json.NewDecoder(r.Body).Decode(&author)

	a, err := models.UpdateAuthor(author)

	if err != nil{
		response := map[string]interface{}{
			"message": "something went wrong",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("author", "UpdateAuthor", "controllers").Inc()
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"author": a,
	}
	jsonResponse, _ := json.Marshal(response)
	//get the remaining timefrom the start and give it to the histogram
	duration := time.Since(start).Seconds()
	prometheus_book.RequestDuration.Observe(duration)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request){
	//timestamp of the handler being called
	start := time.Now()
	
	vars := mux.Vars(r)

	//increment the request counter each time a user make a request call
	prometheus_book.TotalRequest.Inc()

	id,_ := strconv.Atoi(vars["id"])

	msg, err := models.DeleteAuthorById(id)

	if err != nil{
		response := map[string]interface{}{
			"message": "invalid author id",
		}
		//increment the error counter each time a user make a request call and raise an error
		prometheus_book.TotalErros.WithLabelValues("author", "DeleteAuthor", "controllers").Inc()
		jsonResponse, _ := json.Marshal(response)
		utils.JsonWriter(w, http.StatusInternalServerError, jsonResponse)
		return
	}
	response := map[string]interface{}{
		"message": msg,
	}
	jsonResponse, _ := json.Marshal(response)
	//get the remaining timefrom the start and give it to the histogram
	duration := time.Since(start).Seconds()
	prometheus_book.RequestDuration.Observe(duration)
	utils.JsonWriter(w, http.StatusOK, jsonResponse)
}