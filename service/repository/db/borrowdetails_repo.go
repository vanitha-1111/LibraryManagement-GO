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
// 1) decrement book copies safely (only if >0)
// 2) insert borrowdetail record with status 'pending'
// commits only if both succeed
func (r *BorrowDetailRepoImpl) CreateBorrowDetail(bd *models.BorrowDetail) (*models.BorrowDetail, error) {
	tx, err := r.DB.Beginx()
	if err != nil {
		return nil, err
	}
	// if anything goes wrong we rollback
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// Attempt to decrement book copies; if no rows returned -> book unavailable
	var remaining int
	if err = tx.Get(&remaining, DecrementBookCopiesQuery, bd.BookID); err != nil {
		return nil, errors.New("book not available or out of copies")
	}

	// prepare borrow detail record
	bd.BorrowStatus = "pending"
	bd.DateReturn = nil

	// insert borrow detail using named params
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

	// commit transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	bd.BorrowDetailsID = insertedID
	return bd, nil
}

// ReturnBorrowDetail: transactional
// - ensure borrowdetail exists and not already returned
// - set borrow_status = 'returned', date_return = now
// - increment book copies
func (r *BorrowDetailRepoImpl) ReturnBorrowDetail(borrowDetailID int) error {
	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// fetch the borrowdetail row
	var bd models.BorrowDetail
	if err = tx.Get(&bd, GetBorrowDetailByIDQuery, borrowDetailID); err != nil {
		return err
	}

	if bd.BorrowStatus == "returned" {
		return errors.New("borrow detail already returned")
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	// update borrowdetails
	if _, err = tx.Exec(UpdateBorrowDetailStatusQuery, "returned", now, borrowDetailID); err != nil {
		return err
	}

	// increment book copies
	var newCopies int
	if err = tx.Get(&newCopies, IncrementBookCopiesQuery, bd.BookID); err != nil {
		return err
	}

	// commit
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
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
