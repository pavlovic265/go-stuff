package models

import (
	"log"

	"gorm.io/gorm"
)

type Book struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type BookModel struct {
	db *gorm.DB
}

func NewBookModel(db *gorm.DB) *BookModel {
	return &BookModel{db}
}

func (bm *BookModel) CreateBook(book *Book) (*Book, error) {
	log.Println(">>> BookModel:CreateBook")
	if result := bm.db.Create(&book); result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (bm *BookModel) GetAllBooks() ([]Book, error) {
	log.Println(">>> start BookModal:GetAllBooks")
	var books []Book

	if result := bm.db.Find(&books); result.Error != nil {
		return nil, result.Error
	}

	log.Println(">>> end BookModal:GetAllBooks")
	return books, nil
}

func (bm *BookModel) GetBookById(id int64) (*Book, error) {
	var book Book

	if result := bm.db.First(&book, id); result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (bm *BookModel) DeleteBook(id int64) error {
	var book Book
	if result := bm.db.Delete(&book, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (bm *BookModel) UpdateBook(book *Book) error {

	if result := bm.db.Save(&book); result.Error != nil {
		return result.Error
	}

	return nil
}
