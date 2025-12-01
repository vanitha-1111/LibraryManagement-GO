package handler

import (
	"errors"
	"library/service/models"
	"library/service/repository"

	"github.com/gin-gonic/gin"
)

type BorrowDetailService struct {
	borrowDetailRepo repository.BorrowDetailRepo
	borrowRepo       repository.BorrowRepo
	memberRepo       repository.MemberRepo
	bookRepo         repository.BookRepo
}

func NewBorrowDetailService(
	bdRepo repository.BorrowDetailRepo,
	borrowRepo repository.BorrowRepo,
	memberRepo repository.MemberRepo,
	bookRepo repository.BookRepo,
) *BorrowDetailService {
	return &BorrowDetailService{
		borrowDetailRepo: bdRepo,
		borrowRepo:       borrowRepo,
		memberRepo:       memberRepo,
		bookRepo:         bookRepo,
	}
}

// CreateBorrowDetail: validate business rules then call repo (which is transactional)
func (s *BorrowDetailService) CreateBorrowDetail(bd *models.BorrowDetail) (*models.BorrowDetail, error) {
	// 1) check borrow header exists
	borrow, err := s.borrowRepo.GetBorrowByID(bd.BorrowID)
	if err != nil {
		return nil, errors.New("borrow transaction not found")
	}

	// 2) check member not banned
	member, err := s.memberRepo.GetMemberByID(borrow.MemberID)
	if err != nil {
		return nil, errors.New("member not found")
	}
	if member.Status == "Banned" || member.Status == "banned" {
		return nil, errors.New("member is banned")
	}

	// 3) check book exists
	_, err = s.bookRepo.GetBookByID(bd.BookID)
	if err != nil {
		return nil, errors.New("book not found")
	}

	// 4) delegate to repository which does the transaction (decrement + insert)
	return s.borrowDetailRepo.CreateBorrowDetail(bd)
}

// ReturnBorrowDetail: delegate to repo which handles transaction and incrementing
func (s *BorrowDetailService) ReturnBorrowDetail(borrowDetailID int) error {
	return s.borrowDetailRepo.ReturnBorrowDetail(borrowDetailID)
}

// GetBorrowDetailsByBorrowID: delegate to repo
func (s *BorrowDetailService) GetBorrowDetailsByBorrowID(borrowID int) ([]models.BorrowDetail, error) {
	return s.borrowDetailRepo.GetBorrowDetailsByBorrowID(borrowID)
}

func (s *BorrowDetailService) GetMemberBorrowHistory(memberID int) (interface{}, error) {
	// 1) get raw rows
	rows, err := s.borrowDetailRepo.GetMemberBorrowHistory(memberID)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return gin.H{"message": "no history found"}, nil
	}

	// 2) group by borrow_id
	history := make(map[int]gin.H)

	for _, r := range rows {
		if _, exists := history[r.BorrowID]; !exists {
			history[r.BorrowID] = gin.H{
				"borrow_id":   r.BorrowID,
				"date_borrow": r.DateBorrow,
				"due_date":    r.DueDate,
				"books":       []gin.H{},
			}
		}

		bookEntry := gin.H{
			"borrow_details_id": r.BorrowDetailsID,
			"book_id":           r.BookID,
			"book_title":        r.BookTitle,
			"borrow_status":     r.BorrowStatus,
			"date_return":       r.DateReturn,
		}

		history[r.BorrowID]["books"] = append(history[r.BorrowID]["books"].([]gin.H), bookEntry)
	}

	// convert map to slice
	var response []gin.H
	for _, item := range history {
		response = append(response, item)
	}

	return response, nil
}
