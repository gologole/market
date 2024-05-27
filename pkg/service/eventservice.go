package service

import (
	"main.go/models"
	"main.go/pkg/repository"
)

type EventService interface {
	GetEventById(id int) (*models.Hachaton, error)
	GetAllEvents() ([]*models.Hachaton, error)
	CreateEvent(event *models.Hachaton) error
	UpdateEvent(event *models.Hachaton) error
}

type eventService struct {
	repo repository.HachatonRepository
}

func NewEventService(repo repository.HachatonRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) GetEventById(id int) (*models.Hachaton, error) {
	return s.repo.GetEventByID(id)
}

func (s *eventService) GetAllEvents() ([]*models.Hachaton, error) {
	return s.repo.GetAllEvents()
}

func (s *eventService) CreateEvent(event *models.Hachaton) error {
	return s.repo.CreateEvent(event)
}

func (s *eventService) UpdateEvent(event *models.Hachaton) error {
	return s.repo.UpdateEvent(event)
}
