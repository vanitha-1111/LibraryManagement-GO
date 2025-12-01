package handler

import (
	"errors"
	"library/service/models"
	"library/service/repository"
)

type BookService struct {
	bookRepo repository.BookRepo
}

func NewBookService(bookRepo repository.BookRepo) *BookService {
	return &BookService{bookRepo: bookRepo}
}

// CreateBook adds a new book
func (s *BookService) CreateBook(book *models.Book) (int, error) {
	if book.BookTitle == "" {
		return 0, errors.New("Book title is required")
	}
	id, err := s.bookRepo.CreateBook(book)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetBooks
func (s *BookService) ListBooks() ([]models.Book, error) {
	return s.bookRepo.ListBooks()
}
func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.bookRepo.GetBookByID(id)
}

// GetBooksByCategpryName
func (s *BookService) ListBooksByCategoryName(name string) ([]models.Book, error) {
	return s.bookRepo.ListBooksByCategoryName(name)
}
