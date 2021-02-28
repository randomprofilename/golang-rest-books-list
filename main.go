package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := driver.ConnectDB()

	defer db.Close()

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	booksRouter := router.PathPrefix("/books").Subrouter()

	booksController := controllers.NewBooksController(db)
	booksRoutes := routes.NewBooksRoutes(booksController)

	booksRoutes.InitializeBooksRoutes(booksRouter)

	log.Fatal(http.ListenAndServe(":8000", router))
}
