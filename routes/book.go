package routes

import (
	"books-list/controllers"

	"github.com/gorilla/mux"
)

// BooksRoutes describes routes for http requests
type BooksRoutes struct {
	booksController controllers.BooksController
}

// NewBooksRoutes creates and returns new instance of BooksRoutes
func NewBooksRoutes(booksController controllers.BooksController) *BooksRoutes {
	return &BooksRoutes{booksController}
}

// InitializeBooksRoutes initializes routes from BooksController into router from argument
func (b BooksRoutes) InitializeBooksRoutes(router *mux.Router) {
	router.HandleFunc("/", b.booksController.GetBooks).Methods("GET")
	router.HandleFunc("/{id}", b.booksController.GetBook).Methods("GET")
	router.HandleFunc("/", b.booksController.AddBook).Methods("POST")
	router.HandleFunc("/", b.booksController.UpdateBook).Methods("PUT")
	router.HandleFunc("/{id}", b.booksController.RemoveBook).Methods("DELETE")
}
