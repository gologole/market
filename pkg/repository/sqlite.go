package repository

import (
	"database/sql"
	"fmt"
	"log"
	"main.go/models"
	"strings"
)

type Repository struct {
	User     UserRepository
	Team     TeamRepository
	Hachaton HachatonRepository
}

func NewRepository(userdb *sql.DB, teamdb *sql.DB, hackdb *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(userdb),
		Team: NewTeamRepository(teamdb),
		//Hachaton: NewHachatonRepository(hackdb),
	}
}

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	GetUserByLogin(login string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
}

type TeamRepository interface {
	CreateTeam(team *models.Team) error
	GetTeamByID(id int) (*models.Team, error)
	GetAllTeams() ([]*models.Team, error)
	UpdateTeam(team *models.Team) error
	DeleteTeam(id int) error
}

type HachatonRepository interface {
	CreateHachaton(hachaton *models.Hachaton) error
	GetHachatonByID(id int) (*models.Hachaton, error)
	UpdateHachaton(hachaton *models.Hachaton) error
	DeleteHachaton(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	// Логирование значений для проверки
	log.Printf("Creating user: %+v\n", user)

	query := `INSERT INTO users (login, page, password, password_hash, email, role, address, phone_number, description, team_id, skills) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, user.Login, user.Page, user.Password, user.PasswordHash, user.Email, user.Role, user.Address, user.PhoneNumber, user.Description, user.TeamID, user.Skills)
	if err != nil {
		return fmt.Errorf("could not create user: %v", err)
	}
	return nil
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, login, page, password_hash, email, role, address, phone_number, description, team_id, skills 
              FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Login, &user.Page, &user.PasswordHash, &user.Email, &user.Role, &user.Address, &user.PhoneNumber, &user.Description, &user.TeamID, &user.Skills)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get user: %v", err)
	}
	return user, nil
}

func (r *userRepository) GetUserByLogin(login string) (*models.User, error) {
	query := `SELECT id, login, page, password, password_hash, email, role, address, phone_number, description, team_id, skills 
              FROM users WHERE login = ?`
	row := r.db.QueryRow(query, login)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Login, &user.Page, &user.Password, &user.PasswordHash, &user.Email, &user.Role, &user.Address, &user.PhoneNumber, &user.Description, &user.TeamID, &user.Skills)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get user: %v", err)
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET login = ?, page = ?, password_hash = ?, email = ?, role = ?, address = ?, phone_number = ?, description = ?, team_id = ?, skills = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Login, user.Page, user.PasswordHash, user.Email, user.Role, user.Address, user.PhoneNumber, user.Description, user.TeamID, user.Skills, user.ID)
	if err != nil {
		return fmt.Errorf("could not update user: %v", err)
	}
	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete user: %v", err)
	}
	return nil
}
func (r *userRepository) GetAllUsers() ([]*models.User, error) {
	query := `SELECT id, login, page, password, password_hash, email, role, address, phone_number, description, team_id, skills FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Login, &user.Page, &user.Password, &user.PasswordHash, &user.Email, &user.Role, &user.Address, &user.PhoneNumber, &user.Description, &user.TeamID, &user.Skills)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %v", err)
	}
	return users, nil
}

type teamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) *teamRepository {
	return &teamRepository{
		db: db,
	}
}

func (r *teamRepository) CreateTeam(team *models.Team) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	// Insert into teams table
	query := `INSERT INTO teams (name, page, description, is_fully, who_need, hachatons, won, story) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	res, err := tx.Exec(query, team.Name, team.Page, team.Description, team.IsFully, strings.Join(team.WhoNeed, ","), team.Hachatons, team.Won, team.Story)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not create team: %v", err)
	}

	teamID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not get last insert id: %v", err)
	}

	// Insert into teamates table
	for _, teamate := range team.Teamate {
		query = `INSERT INTO teamates (team_id, user_id) VALUES (?, ?)`
		_, err := tx.Exec(query, teamID, teamate.ID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("could not create teamate: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}

func (r *teamRepository) GetTeamByID(id int) (*models.Team, error) {
	query := `SELECT id, name, page, description, is_fully, who_need, hachatons, won, story FROM teams WHERE id = ?`
	row := r.db.QueryRow(query, id)

	team := &models.Team{}
	var whoNeed string
	err := row.Scan(&team.ID, &team.Name, &team.Page, &team.Description, &team.IsFully, &whoNeed, &team.Hachatons, &team.Won, &team.Story)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("could not get team: %v", err)
	}
	team.WhoNeed = strings.Split(whoNeed, ",")
	return team, nil
}

func (r *teamRepository) UpdateTeam(team *models.Team) error {
	query := `UPDATE teams SET name=?, page=?, description=?, is_fully=?, who_need=?, hachatons=?, won=?, story=? WHERE id=?`
	_, err := r.db.Exec(query, team.Name, team.Page, team.Description, team.IsFully, strings.Join(team.WhoNeed, ","), team.Hachatons, team.Won, team.Story, team.ID)
	if err != nil {
		return fmt.Errorf("could not update team: %v", err)
	}
	return nil
}

func (r *teamRepository) DeleteTeam(id int) error {
	query := `DELETE FROM teams WHERE id=?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not delete team: %v", err)
	}
	return nil
}

func (r *teamRepository) GetAllTeams() ([]*models.Team, error) {
	query := `SELECT id, name, page, description, is_fully, who_need, hachatons, won, story FROM teams`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not get teams: %v", err)
	}
	defer rows.Close()

	var teams []*models.Team
	for rows.Next() {
		team := &models.Team{}
		var whoNeed string
		err := rows.Scan(&team.ID, &team.Name, &team.Page, &team.Description, &team.IsFully, &whoNeed, &team.Hachatons, &team.Won, &team.Story)
		if err != nil {
			return nil, fmt.Errorf("could not scan team: %v", err)
		}
		team.WhoNeed = strings.Split(whoNeed, ",")
		teams = append(teams, team)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over teams: %v", err)
	}
	return teams, nil
}

/*
type hachatonRepository struct {
	db *sql.DB
}

func NewHachatonRepository(db *sql.DB) HachatonRepository {
	return &hachatonRepository{db: db}
}

func (r *hachatonRepository) CreateHachaton(hachaton *models.Hachaton) error {
	// Implementation...
}

func (r *hachatonRepository) GetHachatonByID(id int) (*models.Hachaton, error) {
	// Implementation...
}

func (r *hachatonRepository) UpdateHachaton(hachaton *models.Hachaton) error {
	// Implementation...
}

func (r *hachatonRepository) DeleteHachaton(id int) error {
	// Implementation...
}*/
