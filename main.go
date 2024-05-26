package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main.go/pkg/myhandler"
	"main.go/pkg/repository"
	"main.go/pkg/service"
	"main.go/server"
)

func main() {
	//initconfigs()
	port := "8080"

	udb, tdb, edb, err := Migrate("./databases")
	if err != nil {
		log.Fatal("По неествественным причинам не получилось подключиться к бд: ", err)
	}

	r := repository.NewRepository(udb, tdb, edb)

	myservice := service.NewService(*r)

	handler := myhandler.NewMyHandler(myservice)

	server := new(server.Server)
	if err := server.RunServer(port, handler.InitRouts()); err != nil {
		log.Fatal("Server start error: ", err)
	}
	fmt.Scanln()
}

// Migrate подключается к соответствующим базам данных и возвращает дескрипторы или создает их, если таковых не существует.
func Migrate(pathto string) (userDB *sql.DB, teamDB *sql.DB, eventDB *sql.DB, err error) {
	// Подключаемся к базам данных или создаем их, если они не существуют
	userDB, err = sql.Open("sqlite3", fmt.Sprintf("%s/user.db", pathto))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to user database: %v", err)
	}

	teamDB, err = sql.Open("sqlite3", fmt.Sprintf("%s/team.db", pathto))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to team database: %v", err)
	}

	eventDB, err = sql.Open("sqlite3", fmt.Sprintf("%s/event.db", pathto))
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to connect to event database: %v", err)
	}

	// Создаем таблицы, если их не существует
	err = createTables(userDB, teamDB, eventDB)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create tables: %v", err)
	}

	return userDB, teamDB, eventDB, nil
}

// createTables создает необходимые таблицы в базах данных.
func createTables(userDB, teamDB, eventDB *sql.DB) error {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT,
		page TEXT,
		password TEXT,
		password_hash TEXT,
		email TEXT,
		role TEXT,
		address TEXT,
		phone_number TEXT,
		description TEXT,
		team_id INTEGER,
		skills TEXT
	);`

	teamTable := `
	CREATE TABLE IF NOT EXISTS teams (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		page TEXT,
		description TEXT,
		is_fully BOOLEAN,
		who_need TEXT,
		rating TEXT
	);`

	eventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		page TEXT,
		description TEXT,
		participant TEXT
	);`

	_, err := userDB.Exec(userTable)
	if err != nil {
		return fmt.Errorf("failed to create user table: %v", err)
	}

	_, err = teamDB.Exec(teamTable)
	if err != nil {
		return fmt.Errorf("failed to create team table: %v", err)
	}

	_, err = eventDB.Exec(eventTable)
	if err != nil {
		return fmt.Errorf("failed to create event table: %v", err)
	}

	return nil
}
