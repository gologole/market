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

	// Раздача HTML файла из корня проекта
	router.GET("/site", func(c *gin.Context) {
		c.File("index.html")
	})

	router.Group("/user")
	{
		router.Group("/auth")
		{
			router.POST("/login", h.Signin)  //принмает json с login и пароль
			router.POST("/signup", h.Signup) //принимает сереализованного юзера
		}

		router.GET("/profile/:id", h.GetProfileHandler)             //возвращает юзера с заданным id
		router.POST("/getuserbyparams", h.FindUsersBySkillsHandler) //возвращает массив юзеров подходящих по спецификации
		router.GET("/profiles", h.GetProfileListHandler)            //возвращает массив юзеров
		router.DELETE("/profile/:id", h.DeleteProfileHandler)       //принимает только id удаляемого
		router.POST("/profile/updateprofile/:id", h.UpdateUser)     //принимает id и измененную запись (со всеми полями)
		router.GET("/getuserbyteam/:id", h.GetUsersByTeam)
		router.GET("/getsortrating", h.GetSortRating)
	}

	router.Group("team")
	{
		router.GET("getteam/:id", h.GetTeam)
		router.GET("/getteams", h.GetTeams)
		router.POST("/createteam", h.CreateTeam)
		router.PUT("/:id", h.UpdateTeam)
	}

	router.Group("events")
	{
		router.GET("getevent/:id", h.GetEvent)
		router.GET("/getevents", h.GetEvents)
		router.POST("/createEvent", h.CreateEvent)
		router.PUT("updateEvent/:id", h.UpdateEvent)
	}

	return router
}
