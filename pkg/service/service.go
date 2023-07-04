package service

import (
	simplerestapi "github.com/2hard4me/simple-rest-api"
	"github.com/2hard4me/simple-rest-api/pkg/repository"
)

type User interface {
	Create(user simplerestapi.User) (int, error)
	GetAll() ([]simplerestapi.User, error)
	GetById(id int) (simplerestapi.User, error)
	Delete(id int) error
	Update(id int, input simplerestapi.UpdateUserInput) error
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}