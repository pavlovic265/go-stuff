package main

import (
	"log"
	"net/http"

	"example.com/m/v2/pkg/config"
	"example.com/m/v2/pkg/controller"
	"example.com/m/v2/pkg/models"
	"example.com/m/v2/pkg/routes"
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
