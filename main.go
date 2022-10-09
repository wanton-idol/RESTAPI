package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/records", GetRecords).Methods("GET")
	r.HandleFunc("/records/{id}", GetRecord).Methods("GET")
	r.HandleFunc("/records", CreateRecord).Methods("POST")
	r.HandleFunc("/records/{id}", UpdateRecord).Methods("PUT")
	r.HandleFunc("/records/{id}", DeleteRecord).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}
