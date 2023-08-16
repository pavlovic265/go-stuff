package controller

import (
	"encoding/json"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"log"
	"net/http"
)

type BookController struct {
	bookModel *models.BookModel
}

func NewBookController(bookModel *models.BookModel) *BookController {
	return &BookController{bookModel}
}

func (bc *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println(">>> BookController:GetAllBooks")

	books, err := bc.bookModel.GetAllBooks()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	// no need to marshal
	// res, err = json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (bc *BookController) GetBookById(w http.ResponseWriter, r *http.Request) {
	log.Println(">>> BookController:GetBookById")

	bookId, err := utils.ParseBookId(r)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	book, err := bc.bookModel.GetBookById(bookId)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Println(">>> BookController:CreateBook")
	createBook := &models.Book{}

	err := utils.ParseBody(r, createBook)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	book, err := bc.bookModel.CreateBook(createBook)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Println(">>> BookController:UpdateBook")

	bookId, err := utils.ParseBookId(r)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	book, err := bc.bookModel.GetBookById(bookId)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	if err := bc.bookModel.UpdateBook(book); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println(">>> BookController:DeleteBook")
	bookId, err := utils.ParseBookId(r)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	if err := bc.bookModel.DeleteBook(bookId); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookId)

}
