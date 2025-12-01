package repository

import (
	//"LibraryManagement/service/models"
	"library/service/models"
)

type BookRepo interface {
	CreateBook(book *models.Book) (int, error)
	GetBookByID(id int) (*models.Book, error)
	ListBooks() ([]models.Book, error)
	ListBooksByCategoryName(name string) ([]models.Book, error)

	//DeleteById(id int)
}

type MemberRepo interface {
	CreateMember(member *models.Member) (*models.Member, error)
	GetMemberByID(id int) (*models.Member, error)
	ListMembers() ([]models.Member, error)
	UpdateMember(member *models.Member) (*models.Member, error)
	DeleteMember(id int) error
}

type BorrowRepo interface {
	CreateBorrow(borrow *models.Borrow) (int, error)
	GetBorrowByID(id int) (*models.Borrow, error)
	ListBorrows() ([]models.Borrow, error)
}

type BorrowDetailRepo interface {
	CreateBorrowDetail(bd *models.BorrowDetail) (*models.BorrowDetail, error)
	ReturnBorrowDetail(borrowDetailID int) error
	GetBorrowDetailsByBorrowID(borrowID int) ([]models.BorrowDetail, error)
	GetMemberBorrowHistory(memberID int) ([]models.BorrowHistoryItem, error)
}
type CategoryRepo interface {
	GetAllCategories() ([]models.Category, error)
	CreateCategory(name string) (int, error)
	DeleteCategory(id int) error
}
type UserRepo interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(u *models.User) (*models.User, error)
}
