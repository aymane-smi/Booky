package controllers

import (
	"net/http"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
)

r := mux.NewRouter();

func GetBookById(w http.ResponseWriter, r *http.Request){
	utils.JsonHeader()

	vars := mux.vars(r)

	
}