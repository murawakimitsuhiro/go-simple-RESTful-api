package main

import (
	"fmt"

	"github.com/murawakimitsuhiro/go-simple-RESTful-api/models"

	"github.com/jinzhu/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *gorm.DB

const (
	dbName = "go_simple_RESTful"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "33066"
)

func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	dbSourceStr := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dbSourceStr)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.Memo{})
}

func routers() *chi.Mux {
	//router.Get("/posts", AllPosts)
	//router.Get("/posts/{id}", DetailPost)
	//router.Post("/posts", CreatePost)
	//router.Put("/posts/{id}", UpdatePost)
	//router.Delete("/posts/{id}", DeletePost)

	return router
}
