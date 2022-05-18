package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SanskritiHarmukh/Golang-Projects/tree/main/Go-Bookstore/pkg/models"
	"github.com/SanskritiHarmukh/Golang-Projects/tree/main/Go-Bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

func getBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.getAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing..")
	}
	bookDetails, _ := models.getBookById(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.createBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while Parsing..")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book()
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing..")
	}
	bookDetails, db := models.getBookById(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
