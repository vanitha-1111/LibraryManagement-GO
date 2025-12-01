package db

import (
	"library/service/models"

	"github.com/jmoiron/sqlx"
)

type UserRepoImpl struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepoImpl {
	return &UserRepoImpl{DB: db}
}

func (r *UserRepoImpl) GetUserByUsername(username string) (*models.User, error) {
	var u models.User
	err := r.DB.Get(&u, GetUserByUsernameQuery, username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepoImpl) CreateUser(u *models.User) (*models.User, error) {
	rows, err := r.DB.NamedQuery(InsertUserQuery, u)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&u.UserId)
	}
	return u, nil
}
