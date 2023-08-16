package main

import (
	"go-bookstore/pkg/config"
	"go-bookstore/pkg/controller"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db := config.InitDB()
	bm := models.NewBookModel(db)
	bc := controller.NewBookController(bm)

	routes.RegisterBookStoreRouts(r, bc)

	http.Handle("/", r)

	log.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
