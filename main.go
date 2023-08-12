package main

import (
	"fmt"
	"net/http"

	"github.com/aymane-smi/api-test/controllers"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main(){

	r := mux.NewRouter()

	utils.InitLogger()

	//book routes

	r.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")

	r.HandleFunc("/book", controllers.AddBook).Methods("POST")

	r.HandleFunc("/book", controllers.UpdateBook).Methods("PUT")

	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

	//athor routes

	r.HandleFunc("/author/{id}", controllers.GetAuthorById).Methods("GET")

	r.HandleFunc("/author", controllers.AddAuthor).Methods("POST")

	r.HandleFunc("/author", controllers.UpdateAuthor).Methods("PUT")

	r.HandleFunc("/author/{id}", controllers.DeleteAuthor).Methods("DELETE")

	//prometheus route

	http.Handle("/metrics", promhttp.Handler())

	

	fmt.Println("Start listening to server at port 8000")

	http.ListenAndServe(":8000", r)

}