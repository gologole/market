package service

import (
	"main.go/models"
	"main.go/pkg/repository"
)

// Service объединяет все сервисы.
type Service struct {
	AuthService
	//EventService
	TeamService
	UserService
}

// NewService создает новый экземпляр объединенного сервиса.
func NewService(r repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(r.User),
		//EventService: NewEventService(r.Hachaton),
		//TeamService:  NewTeamService(r.Team),
		UserService: NewUserService(r.User),
	}
}

/*
// EventService предоставляет методы для работы с событиями.
type EventService interface {
	// объявления методов...
}

// TeamService предоставляет методы для работы с командами.
type TeamService interface {
	// объявления методов...
}
*/

// UserService предоставляет методы для работы с пользователями.
type UserService interface {
	GetProfileByID(id int) (*models.User, error)
	GetProfileList() []*models.User
	DeleteProfile(id int) error
}
