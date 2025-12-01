package db

import (
	"library/service/models"

	"github.com/jmoiron/sqlx"
)

type BorrowRepoImpl struct {
	DB *sqlx.DB
}

func NewBorrowRep(db *sqlx.DB) *BorrowRepoImpl {
	return &BorrowRepoImpl{DB: db}
}

// Create Borrow
func (r *BorrowRepoImpl) CreateBorrow(b *models.Borrow) (int, error) {
	rows, err := r.DB.NamedQuery(InsertBorrowQuery, b)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var id int
	if rows.Next() {
		rows.Scan(&id)
	}
	return id, nil
}

//Get Borrow ID

func (r *BorrowRepoImpl) GetBorrowByID(id int) (*models.Borrow, error) {
	var borrow models.Borrow
	err := r.DB.Get(&borrow, GetBorrowByIDQuery, id)
	if err != nil {
		return nil, err
	}
	return &borrow, nil
}

// List All Borrows Records
func (r *BorrowRepoImpl) ListBorrows() ([]models.Borrow, error) {
	var list []models.Borrow
	err := r.DB.Select(&list, ListBorrowsQuery)
	if err != nil {
		return nil, err
	}
	return list, nil
}
