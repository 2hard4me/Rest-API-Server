package repository

import (
	simplerestapi "github.com/2hard4me/simple-rest-api"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user simplerestapi.User) (int, error)
	GetAll() ([]simplerestapi.User, error)
	GetById(id int) (simplerestapi.User, error)
	Delete(id int) error
	Update(id int, input simplerestapi.UpdateUserInput) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}