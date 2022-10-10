package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(127.0.0.1:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"

type Record struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Marks     int       `json:"marks"`
	CreatedAt time.Time `json:"created_at"`
}

type Query struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	MinMarks  int    `json:"min_marks"`
	MaxMarks  int    `json:"max_marks"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Record{})
}

// Function to filter using the request payload
func FilterRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var records []Record
	var queries Query
	json.NewDecoder(r.Body).Decode(&queries)
	startDate := queries.StartDate
	endDate := queries.EndDate
	minMarks := queries.MinMarks
	maxMarks := queries.MaxMarks
	DB.Where(DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Where(DB.Where("marks BETWEEN ? AND ?", minMarks, maxMarks))).Find(&records)
	json.NewEncoder(w).Encode(records)
}

func GetRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var records []Record
	DB.Find(&records)
	json.NewEncoder(w).Encode(records)
}

func GetRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var record Record
	DB.First(&record, params["id"])
	json.NewEncoder(w).Encode(record)
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var record Record
	json.NewDecoder(r.Body).Decode(&record)
	DB.Create(&record)
	json.NewEncoder(w).Encode(record)
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var record Record
	DB.First(&record, params["id"])
	json.NewDecoder(r.Body).Decode(&record)
	DB.Save(&record)
	json.NewEncoder(w).Encode(record)
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var record Record
	DB.Delete(&record, params["id"])
	json.NewEncoder(w).Encode("The record is Deleted Successfully!")
}
