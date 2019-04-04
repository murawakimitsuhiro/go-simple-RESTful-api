package main

import (
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
)

var router *chi.Mux
var db *sql.DB

const (
    dbName = "go-mysql-crud"
    dbPass = "12345"
    dbHost = "localhost"
    dbPort = "33066"
)

func routers() *chi.Mux {
    router.Get("/posts", AllPosts)
    router.Get("/posts/{id}", DetailPost)
    router.Post("/posts", CreatePost)
    router.Put("/posts/{id}", UpdatePost)
    router.Delete("/posts/{id}", DeletePost)
    
    return router
}
