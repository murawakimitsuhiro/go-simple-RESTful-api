package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *gorm.DB

const (
	dbName = "go_simple_RESTful"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "3306"
)

func main() {
	router = chi.NewRouter()

	dbSourceStr := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = gorm.Open("mysql", dbSourceStr)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&Note{})

	routers()
	http.ListenAndServe(":8005", Logger())
}

func routers() *chi.Mux {
	//router.Get("/notes", AllNotes)
	//router.Get("/notes/{id}", DetailNote)
	router.Post("/notes", CreateNote)
	router.Put("/notes/{id}", UpdateNote)
	router.Delete("/notes/{id}", DeleteNote)

	return router
}

type Note struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	json.NewDecoder(r.Body).Decode(&note)

	tx := db.Begin()
	defer tx.Rollback()
	tx.Create(&note)
	tx.Commit()

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note Note
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&note)

	tx := db.Begin()
	defer tx.Rollback()

	tx.First(&note, id)
	tx.Update(&note)

	tx.Commit()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	tx := db.Begin()
	defer tx.Rollback()

	tx.Where("ID = ?", id).Delete(Note{})

	tx.Commit()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Logger return log message
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
