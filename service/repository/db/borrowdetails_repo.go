package db

import (
	"errors"
	"time"

	"library/service/models"

	"github.com/jmoiron/sqlx"
)

type BorrowDetailRepoImpl struct {
	DB *sqlx.DB
}

func NewBorrowDetailRepo(db *sqlx.DB) *BorrowDetailRepoImpl {
	return &BorrowDetailRepoImpl{DB: db}
}

// CreateBorrowDetail: transactional
func (r *BorrowDetailRepoImpl) CreateBorrowDetail(bd *models.BorrowDetail) (*models.BorrowDetail, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil, err
	}
	// Always rollback unless commit succeeds
	defer tx.Rollback()

	// 1) Lock the row
	var currentCopies int
	err = tx.Get(&currentCopies, `SELECT book_copies FROM book WHERE book_id=$1 FOR UPDATE`, bd.BookID)
	if err != nil {
		return nil, err
	}

	if currentCopies <= 0 {
		return nil, errors.New("book is out of stock")
	}

	// 2) Decrement safely
	_, err = tx.Exec(DecrementBookCopiesQuery, bd.BookID)
	if err != nil {
		return nil, err
	}

	// 3) Insert borrow detail
	bd.BorrowStatus = "pending"
	bd.DateReturn = nil

	rows, err := tx.NamedQuery(InsertBorrowDetailQuery, bd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var insertedID int
	if rows.Next() {
		if err = rows.Scan(&insertedID); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("failed to insert borrow detail")
	}

	// 4) Commit
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	bd.BorrowDetailsID = insertedID
	return bd, nil
}

// ReturnBorrowDetail: transactional
func (r *BorrowDetailRepoImpl) ReturnBorrowDetail(borrowDetailID int) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	// Always rollback unless committed
	defer tx.Rollback()

	// 1) Get borrow detail
	var bd models.BorrowDetail
	if err = tx.Get(&bd, GetBorrowDetailByIDQuery, borrowDetailID); err != nil {
		return err
	}

	if bd.BorrowStatus == "returned" {
		return errors.New("borrow detail already returned")
	}

	// 2) Lock book row
	var currentCopies int
	err = tx.Get(&currentCopies, `SELECT book_copies FROM book WHERE book_id=$1 FOR UPDATE`, bd.BookID)
	if err != nil {
		return err
	}

	// 3) Update borrow detail
	now := time.Now().Format("2006-01-02 15:04:05")
	if _, err = tx.Exec(UpdateBorrowDetailStatusQuery, "returned", now, borrowDetailID); err != nil {
		return err
	}

	// 4) Increment book copies
	var newCopies int
	if err = tx.Get(&newCopies, IncrementBookCopiesQuery, bd.BookID); err != nil {
		return err
	}

	// 5) Commit
	return tx.Commit()
}

// GetBorrowDetailsByBorrowID returns details with book titles
func (r *BorrowDetailRepoImpl) GetBorrowDetailsByBorrowID(borrowID int) ([]models.BorrowDetail, error) {
	var list []models.BorrowDetail
	if err := r.DB.Select(&list, GetBorrowDetailsByBorrowIDQuery, borrowID); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *BorrowDetailRepoImpl) GetMemberBorrowHistory(memberID int) ([]models.BorrowHistoryItem, error) {
	var list []models.BorrowHistoryItem

	if err := r.DB.Select(&list, MemberBorrowHistoryQuery, memberID); err != nil {
		return nil, err
	}

	return list, nil
}
