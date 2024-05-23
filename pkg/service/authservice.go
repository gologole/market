package service

import (
	"errors"
	"log"
	"main.go/pkg/repository"
)

type Authservice struct {
	db repository.DB
}

func NewAuthservice(db repository.DB) *Authservice {
	return &Authservice{db: db}
}

// логика обработки прав пользователя
func (s *Authservice) Authorization(login string, password string) (bool, error) {
	id, err := s.Signin(login, password)
	if err != nil {
		log.Println("Authservice error : ", err)
		return false, err
	}
	if id == -1 {
		return false, errors.New("Incorrect login or password")
	}
	if id == 0 {
		//здесь я думаю передавать в контекст запроса инфу о том что пользователь адм
		return true, nil
	}

	return false, err
}

func (a *Authservice) Signin(login, password string) (int, error) {
	id, err := a.db.Login(login, password)

	if err != nil {
		return -1, err
	}
	if id == -1 {
		return -1, errors.New("Incorrect login or password")
	} else {
		return id, nil
	}
}

func (a *Authservice) AddProfile(username string, passwordhash string, email string, address string, PhoneNumber string) {
	a.db.AddProfile(username, passwordhash, email, address, PhoneNumber)
}
