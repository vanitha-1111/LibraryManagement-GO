package db

import (
	"errors"
	"library/service/models"

	"github.com/jmoiron/sqlx"
	//"library/service/repository/db"
)

type BookRepoImpl struct {
	DB *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) *BookRepoImpl {
	return &BookRepoImpl{DB: db}
}
func (r *BookRepoImpl) CreateBook(book *models.Book) (int, error) {
	rows, err := r.DB.NamedQuery(InsertBookQuery, book)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, errors.New("No Id returned")
}

func (r *BookRepoImpl) GetBookByID(id int) (*models.Book, error) {
	var b models.Book
	if err := r.DB.Get(&b, GetBookByIDQuery, id); err != nil {
		return nil, err
	}
	return &b, nil
}
func (r *BookRepoImpl) ListBooks() ([]models.Book, error) {
	var list []models.Book
	if err := r.DB.Select(&list, ListBooksQuery); err != nil {
		return nil, err
	}
	return list, nil
}
func (r *BookRepoImpl) ListBooksByCategoryName(name string) ([]models.Book, error) {
	var list []models.Book
	if err := r.DB.Select(&list, ListBooksByCategoryNameQuery, name); err != nil {
		return nil, err
	}
	return list, nil
}
