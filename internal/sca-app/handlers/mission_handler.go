package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"github.com/rvxt21/sca-agency/internal/sca-app/service"
)

type MissionHandler struct {
	s service.MissionService
}

func NewMissionHandler(s service.MissionService) *MissionHandler {
	return &MissionHandler{
		s: s,
	}
}

func (h *MissionHandler) CreateMission(c *gin.Context) {
	var mission models.Mission

	if err := c.ShouldBindJSON(&mission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.s.CreateMission(&mission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create mission"})
		return
	}

	c.JSON(http.StatusCreated, mission)
}
