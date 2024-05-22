package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type DB interface {
	ConnectDB() (*SQLiteDB, error)
	CreateTable()
	AddProfile(username string, passwordhash string, email string, address string, PhoneNumber string)
	Login(username string, passwordhash string) (int, error)
	//DeleteProfile(db *sql.DB, id int)
	//	GetProfileList(db *sql.DB, id int) ([]models.User, error)
	//UpdatePassword(db *sql.DB, id int, password string)
}

type SQLiteDB struct {
	db *sql.DB
}

func (s *SQLiteDB) ConnectDB() (*SQLiteDB, error) {
	_, err := os.Stat("./databases/market.db")
	if os.IsNotExist(err) {
		// Если файл базы данных не существует, создаем новый файл
		file, err := os.Create("./databases/market.db")
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	descriptor, err := sql.Open("sqlite3", "./databases/market.db")
	if err != nil {
		return nil, err
	}

	s.db = descriptor
	return s, nil
}

func (db *SQLiteDB) CreateTable() {
	statement, _ := db.db.Prepare("CREATE TABLE IF NOT EXISTS " +
		"users(id INTEGER)," +
		"login VARCHAR(255)," +
		"passwordhash VARCHAR(255)," +
		"email VARCHAR(255)," +
		"role VARCHAR(255)," +
		"address VARCHAR(255)," +
		"phone_number VARCHAR(255))")
	statement.Exec()
}

func (db *SQLiteDB) AddProfile(username string, passwordhash string, email string, address string, PhoneNumber string) {
	statement, _ := db.db.Prepare("INSERT INTO users(login, passwordhash, email, role, address, phone_number) " +
		"VALUES(?,?,?,?,?,?)")
	statement.Exec(username, passwordhash, email, "user", address, PhoneNumber)
}

// возвращает индекс пользователя по логину или -1, если пользователь не найден
func (db *SQLiteDB) Login(username string, passwordhash string) (int, error) {
	statement, _ := db.db.Prepare("SELECT * FROM users WHERE login =? AND passwordhash =?")
	rows, _ := statement.Query(username, passwordhash)
	defer rows.Close()
	for rows.Next() {
		var id int
		var login string
		var passwordhash string
		var email string
		var role string
		var address string
		var phone_number string
		rows.Scan(&id, &login, &passwordhash, &email, &role, &address, &phone_number)
		return id, nil
	}
	return -1, nil
}
