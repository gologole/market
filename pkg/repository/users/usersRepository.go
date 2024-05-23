package users

import (
	"database/sql"
	"main.go/models"
)

// операции с хранилищем пользователей
type UserRepository interface {
	ConnectDB(*sql.DB)
	CreateTable()
	AddProfile()
	Login()
	DeleteProfile(db *sql.DB, id int)
	GetProfileList(db *sql.DB, id int) ([]models.User, error)
	UpdatePassword(db *sql.DB, id int, password string)
}

type UserRepository struct {
	db *sql.DB
}
