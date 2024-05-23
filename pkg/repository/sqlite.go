package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	ConnectDB() (*SQLiteDB, error)
	CreateTable()
	AddProfile(username string, passwordhash string, email string, address string, PhoneNumber string)
	Login(username string, passwordhash string) (int, error)
	//DeleteProfile(db *sql.DB, id int)
	//GetProfileList(db *sql.DB, id int) ([]models.User, error)
	//UpdatePassword(db *sql.DB, id int, password string)
}

type SQLiteDB struct {
	db *sql.DB
}
