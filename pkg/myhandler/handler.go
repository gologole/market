package myhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/service"
)

type MyHandler struct {
	Service *service.Service
}

func NewMyHandler(service *service.Service) *MyHandler {
	return &MyHandler{
		Service: service,
	}
}

func (h *MyHandler) InitRouts() *gin.Engine {
	router := gin.Default()

	router.Group("/user")
	{
		router.Group("/auth")
		{
			router.POST("/login", h.Signin)  //принмает json с login и пароль
			router.POST("/signup", h.Signup) //принимает сереализованного юзера
		}

		router.GET("/profile/:id", h.GetProfileHandler)         //возвращает юзера с заданным id
		router.GET("/profiles", h.GetProfileListHandler)        //возвращает массив юзеров
		router.DELETE("/profile/:id", h.DeleteProfileHandler)   //принимает только id удаляемого
		router.POST("/profile/updateprofile/:id", h.UpdateUser) //принимает id и измененную запись (со всеми полями)
	}

	router.Group("team")
	{
		router.GET("/:id", h.GetTeam)
		router.GET("/getteams", h.GetTeams)
		router.POST("/createteam", h.CreateTeam)
		router.PUT("/:id", h.UpdateTeam)
	}

	router.Group("events")
	{
		router.GET("/:id", GetEvent)
		router.GET("/getEvent", GetEvents)
		router.POST("/createEvent", CreateEvent)
		router.PUT("/:id", UpdateEvent) //включая запись команд
	}

	return router
}
