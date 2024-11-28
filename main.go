package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/izzuddinm/go-restapi-mux/controllers/productcontroller"
	"github.com/izzuddinm/go-restapi-mux/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/products/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/products", productcontroller.Create).Methods("POST")
	r.HandleFunc("/products/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/products/{id}", productcontroller.Index).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
