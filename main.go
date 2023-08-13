package main

import (
	"fmt"
	"net/http"

	"github.com/aymane-smi/api-test/controllers"
	prometheus_book "github.com/aymane-smi/api-test/prometheus"
	"github.com/aymane-smi/api-test/utils"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
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

	//prometheus route & config

	r.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(prometheus_book.TotalRequest)
	prometheus.MustRegister(prometheus_book.TotalErros)
	prometheus.MustRegister(prometheus_book.RequestDuration)

	

	fmt.Println("Start listening to server at port 8000")

	http.ListenAndServe(":8000", r)

}