package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/handler"

	"github.com/murawakimitsuhiro/go-simple-RESTful-api/driver"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName = "go_simple_RESTful"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "3306"
)

func main() {
	connection, err := driver.ConnectGorm(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(-1)
	}

	r := routers(connection)
	fmt.Println("Server listen at :8005")
	if err := http.ListenAndServe(":8005", r); err != nil {
		fmt.Println(err)
	}
}

func routers(db *driver.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	nHandler := handler.NewNoteHandler(db)

	r.Get("/notes", nHandler.Fetch)
	r.Get("/notes/{id}", nHandler.GetByID)
	r.Post("/notes", nHandler.Create)
	r.Put("/notes/{id}", nHandler.Update)
	r.Delete("/notes/{id}", nHandler.Delete)

	return r
}
