package models

import (
	"github.com/SanskritiHarmukh/Golang-Projects/tree/main/Go-Bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book {
	// db helps us to talk to database
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func getBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id)
	return &getBook, db
}

func deleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
