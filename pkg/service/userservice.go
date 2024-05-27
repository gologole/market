package service

import (
	"fmt"
	"main.go/models"
	"main.go/pkg/repository"
	"sort"
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

func (s *UserServiceStruct) FindUsersBySkills(skills []string) ([]*models.User, error) {
	return s.r.FindUsersBySkills(skills)
}

func (s *UserServiceStruct) GetUsersByTeamID(teamID int) ([]*models.User, error) {
	users, err := s.r.GetUsersByTeamID(teamID)
	if err != nil {
		return nil, fmt.Errorf("could not get users by team ID: %v", err)
	}
	return users, nil
}

func (s *UserServiceStruct) GetUsersSortedByHackatonsAndWon() ([]*models.User, error) {
	users, err := s.r.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Сортировка пользователей по убыванию хакатонов и выигрышей
	sort.Slice(users, func(i, j int) bool {
		if users[i].Hachatons != users[j].Hachatons {
			return users[i].Hachatons > users[j].Hachatons
		}
		return users[i].Won > users[j].Won
	})

	return users, nil
}
