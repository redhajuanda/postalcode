package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/address/{postal_code}", singleAddress).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	handleRequests()
}
