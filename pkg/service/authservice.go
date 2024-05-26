package service

import (
	"errors"
	"fmt"
	"main.go/models"
	"main.go/pkg/repository"
)

type AuthService interface {
	Register(user *models.User) error
	Login(login, password string) (*models.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(user *models.User) error {
	return s.userRepository.CreateUser(user)
}

func (s *authService) Login(login, password string) (*models.User, error) {
	user, err := s.userRepository.GetUserByLogin(login)
	if err != nil {
		return nil, fmt.Errorf("could not find user: %v", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	// Сравниваем незашифрованный пароль
	fmt.Println(user.Password, " и ", "password")
	if user.Password != password {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
