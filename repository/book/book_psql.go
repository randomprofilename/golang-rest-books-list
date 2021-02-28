package bookrepository

import (
	"books-list/models"
	"database/sql"
	"errors"
)

// BookRepository declares method for access books-storage
type BookRepository interface {
	GetBooks() (*[]models.Book, error)
	GetBook(id int) (*models.Book, error)
	AddBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	RemoveBook(id int) error
}

type bookRepository struct {
	db *sql.DB
}

// ErrBookNotExist will be returned when get by id, update or delete methods trying to be applied to non-existent book
var ErrBookNotExist = errors.New("Book isn't exist")

// NewBookRepository is creates new BookRepository instance
func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db}
}

func (b bookRepository) GetBooks() (*[]models.Book, error) {
	rows, err := b.db.Query("select * from books")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	books := []models.Book{}

	for rows.Next() {
		var book models.Book

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return &books, err
}

func (b bookRepository) GetBook(id int) (*models.Book, error) {
	var book models.Book

	err := b.db.QueryRow(
		"select * from books where id=$1", id,
	).Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err == sql.ErrNoRows {
		return nil, ErrBookNotExist
	}

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b bookRepository) AddBook(book *models.Book) error {
	err := b.db.QueryRow(
		"insert into books(title, author, year) values ($1, $2, $3) returning id;",
		book.Title, book.Author, book.Year,
	).Scan(&book.ID)

	if err != nil {
		return err
	}
	return nil
}

func (b bookRepository) UpdateBook(book *models.Book) error {
	result, err := b.db.Exec("update books set title=$1, author=$2, year=$3 where id=$4;",
		book.Title,
		book.Author,
		book.Year,
		book.ID,
	)

	if err == sql.ErrNoRows {
		return ErrBookNotExist
	}

	if err != nil {
		return err
	}

	booksUpdated, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if booksUpdated == 0 {
		return errors.New("Book with this id do not exist")
	}

	return nil
}

func (b bookRepository) RemoveBook(id int) error {
	result, err := b.db.Exec("delete from books where id =$1", id)

	if err != nil {
		return err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsDeleted == 0 {
		return ErrBookNotExist
	}

	return nil
}
