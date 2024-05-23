package users

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/service"
)

type MyHandler struct {
	Service *service.Authservice
}

func (h *MyHandler) InitRouts() *gin.Engine {
	router := gin.Default()

	router.Group("/user")
	{
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
	}

	router.Group("events")
	{
		router.GET("/:id", GetEvent)
		router.GET("/getEvent", GetEvents)
		router.POST("/createEvent", CreateEvent)
		router.PUT("/:id", UpdateEvent)    //включая запись команд
		router.DELETE("/:id", DeleteEvent) //требует авторизации
	}

	router.Group("team")
	{
		router.GET("/:id", GetTeam)
		router.GET("/getTeams", GetTeams)
		router.POST("/getTeamsbyparams", GetTeamsbyparams)
		router.GET("/getRatingList", GetRatingList)
		router.POST("/createTeam", CreateTeam)
		router.PUT("/:id", UpdateTeam)
		router.DELETE("/:id", DeleteTeam) //требует авторизации
	}

	return router
}
