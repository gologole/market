package service

type Service interface {
	Authservice
	TeamService
	EventService
}

type Service struct {
	Authservice
	TeamService
	EventService
}

func NewService() *Service {
	return &Service{
		Authservice:  NewAuthservice(),
		TeamService:  NewTeamService(),
		EventService: NewEventService(),
	}
}
