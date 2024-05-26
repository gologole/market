package myhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/models"
	"net/http"
	"strconv"
)

func (h *MyHandler) GetTeam(c *gin.Context) {
	teamIDStr := c.Param("id")
	teamID, err := strconv.Atoi(teamIDStr)
	if err != nil {
		fmt.Println("GetTeam error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team ID"})
		return
	}

	team, err := h.Service.GetTeamById(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if team == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}

// GetTeams обрабатывает запрос на получение списка всех команд.
func (h *MyHandler) GetTeams(c *gin.Context) {
	teams, err := h.Service.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (h *MyHandler) CreateTeam(c *gin.Context) {
	if h.Service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "service not initialized"})
		return
	}

	var team models.Team
	if err := c.BindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateTeam(&team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Team created successfully"})
}

// UpdateTeam обрабатывает запрос на обновление данных о команде.
func (h *MyHandler) UpdateTeam(c *gin.Context) {
	// Извлекаем ID команды из параметра маршрута
	teamIDStr := c.Param("id")
	teamID, err := strconv.Atoi(teamIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team ID"})
		return
	}

	// Получаем данные о команде из запроса
	var team models.Team
	if err := c.BindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	team.ID = teamID

	// Обновляем данные о команде
	if err := h.Service.UpdateTeam(&team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team updated successfully"})
}
