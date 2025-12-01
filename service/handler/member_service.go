package handler

import (
	"errors"
	"library/service/models"
	"library/service/repository"
)

type MemberService struct {
	memberRepo repository.MemberRepo
}

func NewMemberService(repo repository.MemberRepo) *MemberService {
	return &MemberService{memberRepo: repo}
}

// CreateMember returns the FULL created member with the generated member_id
func (s *MemberService) CreateMember(member *models.Member) (*models.Member, error) {
	if member.Firstname == "" {
		return nil, errors.New("firstname is required")
	}

	createdMember, err := s.memberRepo.CreateMember(member)
	if err != nil {
		return nil, err
	}

	return createdMember, nil
}

func (s *MemberService) ListMembers() ([]models.Member, error) {
	return s.memberRepo.ListMembers()
}

func (s *MemberService) GetMemberByID(id int) (*models.Member, error) {
	return s.memberRepo.GetMemberByID(id)
}
func (s *MemberService) UpdateMember(member *models.Member) (*models.Member, error) {
	if member.MemberID == 0 {
		return nil, errors.New("member_id is required")
	}
	return s.memberRepo.UpdateMember(member)
}

func (s *MemberService) DeleteMember(id int) error {
	return s.memberRepo.DeleteMember(id)
}
