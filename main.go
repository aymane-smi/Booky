package main

import (
	"fmt"
	"net/http"

	"github.com/aymane-smi/api-test/controllers"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
)

func main(){

	r := mux.NewRouter()

	utils.InitLogger()

	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")

	r.HandleFunc("/book", controllers.AddBook).Methods("POST")

	r.HandleFunc("/book", controllers.UpdateBook).Methods("PUT")

	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

	fmt.Println("Start listening to server at port 8000")

	http.ListenAndServe(":8000", r)
}