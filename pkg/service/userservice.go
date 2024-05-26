package service

import (
	"main.go/models"
	"main.go/pkg/repository"
)

type UserServiceStruct struct {
	r repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceStruct{
		userRepository,
	}
}

func (s *UserServiceStruct) GetProfileByID(id int) (*models.User, error) {
	return s.r.GetUserByID(id)
}

func (s *UserServiceStruct) GetProfileList() []*models.User {
	users, _ := s.r.GetAllUsers()
	return users
}

func (s *UserServiceStruct) DeleteProfile(id int) error {
	return s.r.DeleteUser(id)
}

func (s *UserServiceStruct) UpdateUser(user *models.User) error {
	return s.r.UpdateUser(user)
}
