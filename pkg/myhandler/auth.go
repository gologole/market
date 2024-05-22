package myhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"main.go/models"
	"net/http"
)

func (h MyHandler) Signin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// Используем полученные значения логина и пароля
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	id, err := h.Service.Signin(username, password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		fmt.Println("Login user by this id :", id)

		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func (h MyHandler) Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println("Unmarshal signup error : ", err)
	}
	h.Service.AddProfile(user.Login, user.PasswordHash, user.Email, user.Address, user.PhoneNumber)
}
