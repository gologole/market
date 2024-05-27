package myhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/models"
	"net/http"
	"strconv"
)

// GetProfileHandler возвращает профиль пользователя по его ID.
func (h *MyHandler) GetProfileHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.Service.GetProfileByID(id)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// GetProfileListHandler возвращает список всех профилей пользователей.
func (h *MyHandler) GetProfileListHandler(c *gin.Context) {
	users := h.Service.UserService.GetProfileList()
	c.JSON(http.StatusOK, users)
}

// DeleteProfileHandler удаляет профиль пользователя по его ID.
func (h *MyHandler) DeleteProfileHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.Service.UserService.DeleteProfile(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully"})
}
func (h *MyHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Извлекаем ID пользователя из параметра маршрута
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	user.ID = userID

	// Обновляем данные о пользователе
	if err := h.Service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *MyHandler) FindUsersBySkillsHandler(c *gin.Context) {
	var requestBody struct {
		Skills []string `json:"skills"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if len(requestBody.Skills) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing skills parameter"})
		return
	}

	users, err := h.Service.FindUsersBySkills(requestBody.Skills)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *MyHandler) GetUsersByTeam(c *gin.Context) {
	teamIDStr := c.Param("id")
	teamID, err := strconv.Atoi(teamIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid team ID"})
		return
	}

	users, err := h.Service.GetUsersByTeamID(teamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *MyHandler) GetSortRating(c *gin.Context) {
	users, err := h.Service.GetUsersSortedByHackatonsAndWon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
