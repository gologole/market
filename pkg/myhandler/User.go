package myhandler

import (
	"github.com/gin-gonic/gin"
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
