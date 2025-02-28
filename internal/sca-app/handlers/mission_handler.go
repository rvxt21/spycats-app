package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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

func (h *MissionHandler) DeleteMission(c *gin.Context) {
	id := c.Value("id").(uint)

	if err := h.s.DeleteMission(id); err != nil {
		log.Error().Err(err).Msg("Failed to delete mission")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, "Mission deleted")
}

type UpdateMissionStatusReqBody struct {
	isCompleted bool `json:"is_complited"`
}

func (h *MissionHandler) UpdateMissionStatus(c *gin.Context) {
	id := c.Value("id").(uint)
	var req UpdateMissionStatusReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.s.UpdateMissionStatus(id, req.isCompleted); err != nil {
		log.Error().Err(err).Msg("Failed to update mission status")
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, "Mission status updated")
}

func (h *MissionHandler) GetAllMissions(c *gin.Context) {
	missions, err := h.s.GetAllMissions()

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve all missions")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve missions"})
		return
	}

	c.JSON(http.StatusOK, missions)
}

func (h *MissionHandler) GetMission(c *gin.Context) {
	id := c.Value("id").(uint)

	cat, err := h.s.GetMission(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve spy cat by ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve spy cat"})
		return
	}

	c.JSON(http.StatusOK, cat)
}