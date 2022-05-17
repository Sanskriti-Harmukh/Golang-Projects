package routes

import (
	"github.com/SanskritiHarmukh/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookId).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.CreateBook).Methods("POST")
}