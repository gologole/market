package teams

import (
	"database/sql"
	"main.go/models"
)

//операции с хранилищем команд

type TeamRepository struct {
	sql sql.DB
}

type TeamRepository interface {
	GetTeamById(id string) error
	GetTeamList() ([]models.Team, error)
	AddTeam(team models.Team) error
	UpdateTeam(team models.Team) error
}
