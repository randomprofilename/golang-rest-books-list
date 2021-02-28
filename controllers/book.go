package controllers

import (
	"books-list/models"
	bookRepository "books-list/repository/book"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// BooksController defines handlers for http requests
type BooksController interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	RemoveBook(w http.ResponseWriter, r *http.Request)
}

type booksController struct {
	booksRepository bookRepository.BookRepository
}

func NewBooksController(db *sql.DB) BooksController {
	booksRepository := bookRepository.NewBookRepository(db)
	return &booksController{booksRepository: booksRepository}
}

func (c booksController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.booksRepository.GetBooks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (c booksController) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := c.booksRepository.GetBook(id)

	if err == bookRepository.ErrBookNotExist {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (c booksController) AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	err := c.booksRepository.AddBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (c booksController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	err := c.booksRepository.UpdateBook(&book)

	if err == bookRepository.ErrBookNotExist {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (c booksController) RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err == bookRepository.ErrBookNotExist {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.booksRepository.RemoveBook(id)
	json.NewEncoder(w).Encode(id)
}
