package handler

import (
	"errors"

	"library/service/models"
	"library/service/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	u, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err == nil {
		return u, nil
	}
	if u.Password == password {
		return u, nil
	}
	return nil, errors.New("invalid password")
}

// ADMIN
func (s *UserService) CreateUser(u *models.User, rawPassword string) (*models.User, error) {
	//Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hash)
	created, err := s.userRepo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return created, nil
}
