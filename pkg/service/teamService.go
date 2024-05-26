package service

import (
	"main.go/models"
	"main.go/pkg/repository"
)

type TeamService interface {
	GetTeamById(id int) (*models.Team, error)
	CreateTeam(team *models.Team) error
	GetAllTeams() ([]*models.Team, error)
	UpdateTeam(team *models.Team) error
}

type TeamServiceStruct struct {
	TeamRepository repository.TeamRepository
}

func NewTeamService(teamRepository repository.TeamRepository) *TeamServiceStruct {
	return &TeamServiceStruct{
		teamRepository,
	}
}

func (s *TeamServiceStruct) GetTeamById(id int) (*models.Team, error) {
	return s.TeamRepository.GetTeamByID(id)
}

func (s *TeamServiceStruct) CreateTeam(team *models.Team) error {
	return s.TeamRepository.CreateTeam(team)
}

func (s *TeamServiceStruct) GetAllTeams() ([]*models.Team, error) {
	return s.TeamRepository.GetAllTeams()
}

func (s *TeamServiceStruct) UpdateTeam(team *models.Team) error {
	return s.TeamRepository.UpdateTeam(team)
}
