package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wanton-idol/RESTAPI/user"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/records", user.GetRecords).Methods("GET")
	r.HandleFunc("/records/{id}", user.GetRecord).Methods("GET")
	r.HandleFunc("/filterrecords", user.FilterRecords).Methods("GET")
	r.HandleFunc("/records", user.CreateRecord).Methods("POST")
	r.HandleFunc("/records/{id}", user.UpdateRecord).Methods("PUT")
	r.HandleFunc("/records/{id}", user.DeleteRecord).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	fmt.Println("Building RESTful API")
	fmt.Println("API is started...")
	user.InitialMigration()
	initializeRouter()
}
