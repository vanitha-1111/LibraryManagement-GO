package db

import (
	"library/service/models"
	"library/service/repository"

	"github.com/jmoiron/sqlx"
)

type CategoryRepoImpl struct {
	DB *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) repository.CategoryRepo {
	return &CategoryRepoImpl{DB: db}
}
func (r *CategoryRepoImpl) CreateCategory(name string) (int, error) {
	var id int
	err := r.DB.QueryRow(InsertCategoryQuery, name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
func (r *CategoryRepoImpl) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := r.DB.Select(&categories, ListCategoriesQuery); err != nil {
		return nil, err
	}
	return categories, nil
}
func (r *CategoryRepoImpl) DeleteCategory(id int) error {
	_, err := r.DB.Exec(DeleteCategoryQuery, id)
	return err
}
