package service

import (
	"errors"

	"myapp/internal/model"
	"myapp/internal/repository"
)

// UserService holds business logic. It knows nothing about HTTP
// or SQL — it only depends on the repository interface.
type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req model.CreateUserRequest) (model.User, error) {
	if req.Name == "" || req.Email == "" {
		return model.User{}, errors.New("name and email are required")
	}

	user := model.User{
		Name:  req.Name,
		Email: req.Email,
	}
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id int) (model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) ListUsers() ([]model.User, error) {
	return s.repo.List()
}
