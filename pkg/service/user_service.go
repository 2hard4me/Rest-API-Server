package service

import (
	simplerestapi "github.com/2hard4me/simple-rest-api"
	"github.com/2hard4me/simple-rest-api/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user simplerestapi.User) (int, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetAll() ([]simplerestapi.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id int) (simplerestapi.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *UserService) Update(id int, input simplerestapi.UpdateUserInput) error {
	return s.repo.Update(id, input)
}


