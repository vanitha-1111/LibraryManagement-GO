package handler

import (
	"errors"
	"library/service/models"
	"library/service/repository"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(repo repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

// Create Category
func (s *CategoryService) CreateCategory(name string) (*models.Category, error) {
	if name == "" {
		return nil, errors.New("category name is required")
	}
	id, err := s.repo.CreateCategory(name)
	if err != nil {
		return nil, err
	}
	return &models.Category{
		CategoryID: id,
		ClassName:  name,
	}, nil

}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}
func (s *CategoryService) DeleteCategory(id int) error {
	return s.repo.DeleteCategory(id)
}
