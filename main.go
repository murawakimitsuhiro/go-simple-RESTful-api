package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/driver"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/handler"
)

func main() {
	// config := driver.DBConnectionConfig{}
	// err := envconfig.Process("go_simple_restful", &config)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	os.Exit(-1)
	// }

	//connection, err := driver.ConnectGorm(config)
	//if err != nil {
	//	fmt.Printf("Failed to connect to database: %v\n", err)
	//	os.Exit(-1)
	//}

	// r := routers(connection)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/ping", ping)

	fmt.Println("Server listen at :8005")
	if err := http.ListenAndServe(":8005", r); err != nil {
		fmt.Println(err)
	}
}

func routers(db *driver.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/ping", ping)

	nHandler := handler.NewNoteHandler(db)
	r.Get("/notes", nHandler.Fetch)
	r.Get("/notes/{id}", nHandler.GetByID)
	r.Post("/notes", nHandler.Create)
	r.Put("/notes/{id}", nHandler.Update)
	r.Delete("/notes/{id}", nHandler.Delete)

	return r
}

func ping(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(map[string]string{"ping": "pong"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
