package repository

import (
	"database/sql"
	"fmt"
	"log"
	"main.go/models"
)

type Repository struct {
	User     UserRepository
	Team     TeamRepository
	Hachaton HachatonRepository
}

func NewRepository(userdb *sql.DB, teamdb *sql.DB, hackdb *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(userdb),
		//Team:     NewTeamRepository(teamdb),
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

/*
type teamRepository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) CreateTeam(team *models.Team) error {
	// Implementation...
}

func (r *teamRepository) GetTeamByID(id int) (*models.Team, error) {
	// Implementation...
}

func (r *teamRepository) UpdateTeam(team *models.Team) error {
	// Implementation...
}

func (r *teamRepository) DeleteTeam(id int) error {
	// Implementation...
}

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
