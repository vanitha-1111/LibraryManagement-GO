package handler

import (
	"errors"
	"library/service/models"
	"library/service/repository"
)

type BorrowService struct {
	borrowRepo repository.BorrowRepo
}

func NewBorrowService(repo repository.BorrowRepo) *BorrowService {
	return &BorrowService{borrowRepo: repo}
}

func (s *BorrowService) CreateBorrow(b *models.Borrow) (int, error) {
	if b.MemberID == 0 {
		return 0, errors.New("member_id is required")
	}
	return s.borrowRepo.CreateBorrow(b)
}

func (s *BorrowService) GetBorrowByID(id int) (*models.Borrow, error) {
	return s.borrowRepo.GetBorrowByID(id)
}

func (s *BorrowService) ListBorrows() ([]models.Borrow, error) {
	return s.borrowRepo.ListBorrows()
}
