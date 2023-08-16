package routes

import (
	"go-bookstore/pkg/controller"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRouts = func(router *mux.Router, controller *controller.BookController) {
	router.HandleFunc("/books", controller.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", controller.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controller.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controller.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{bookId}", controller.DeleteBook).Methods(http.MethodDelete)

}
