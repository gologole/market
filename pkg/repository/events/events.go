package events

import "database/sql"

// операции с хранилищем хакатонов
type Event struct {
	db *sql.DB
}

type EventRepository interface {
	CreateEvent(event Event) error
	DeleteEvent(event Event) error
	UpdateEvent(event Event) error
	GetEvent(id int) (Event, error)
	GetAllEvents() ([]Event, error)
}
