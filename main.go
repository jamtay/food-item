package main

import (
	"github.com/gorilla/mux"
	"log"
	"food-item-api/resources"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	setUp(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setUp(router *mux.Router) {
	router.HandleFunc("/food-items", resources.GetFoodItems).Methods("GET")
	router.HandleFunc("/food-items", resources.CreateFoodItem).Methods("POST")
	router.HandleFunc("/food-items/{id}", resources.GetFoodItem).Methods("GET")
}