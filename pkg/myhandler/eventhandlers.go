package myhandler

import (
	"github.com/gin-gonic/gin"
	"main.go/models"
	"net/http"
	"strconv"
)

func (h *MyHandler) GetEvent(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	event, err := h.Service.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *MyHandler) GetEvents(c *gin.Context) {
	events, err := h.Service.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}

func (h *MyHandler) CreateEvent(c *gin.Context) {
	var event models.Hachaton
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

func (h *MyHandler) UpdateEvent(c *gin.Context) {
	eventIDStr := c.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	var event models.Hachaton
	if err := c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.Id = eventID

	if err := h.Service.UpdateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
