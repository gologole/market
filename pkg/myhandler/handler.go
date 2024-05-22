package myhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/service"
)

type MyHandler struct {
	Service *service.Authservice
}

func (h *MyHandler) InitRouts() *gin.Engine {
	router := gin.Default()
	router.Group("/auth")
	{
		router.POST("/login", h.Signin)
		router.POST("/signup", h.Signup)
	}

	router.GET("/profile/:id", GetProfile)
	router.GET("/profiles", GetProfileList)
	router.DELETE("/profile/:id", DeleteProfile)

	router.Group("update")
	{
		router.PUT("/profile/:id/password", UpdatePassword)
		router.PUT("/profile/:id/email", UpdateEmail)
		router.PUT("/profile/:id/phone", UpdatePhone)
		router.PUT("/profile/:id/address", UpdateAddress)
		router.PUT("/profile/:id/city", UpdateCity)
		router.PUT("/profile/:id/state", UpdateState)
		router.PUT("/profile/:id/country", UpdateCountry)
	}
	return router
}
