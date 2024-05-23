package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"main.go/pkg/repository/events"
	"main.go/pkg/repository/teams"
	"main.go/pkg/repository/users"
)

type DB interface {
	ConnectToDatabase()
	users.UserRepository
	teams.TeamRepository
	events.EventRepository
}

type SQLiteDB struct {
	db *sql.DB
}
