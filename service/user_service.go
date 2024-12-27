package service

import (
	"errors"
	"microService/model"
	"microService/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
	if user.Name == "" || user.Email == "" {
		return model.User{}, errors.New("invalid user data")
	}
	createdUser, err := s.repo.Save(user)
	if err != nil {
		return model.User{}, err
	}
	return createdUser, nil
}
