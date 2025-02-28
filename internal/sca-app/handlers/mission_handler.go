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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, "Mission deleted")
}

type UpdateMissionStatusReqBody struct {
	IsCompleted bool `json:"is_completed"`
}

func (h *MissionHandler) UpdateMissionStatus(c *gin.Context) {
	id := c.Value("id").(uint)
	var req UpdateMissionStatusReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.s.UpdateMissionStatus(id, req.IsCompleted); err != nil {
		log.Error().Err(err).Msg("Failed to update mission status")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

type AssignCatForMission struct {
	CatId uint `json:"cat_id"`
}

func (h *MissionHandler) SetCatForMission(c *gin.Context) {
	id := c.Value("id").(uint)

	var req AssignCatForMission

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.CatId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cat ID"})
		return
	}

	if err := h.s.AssignCatToMission(id, req.CatId); err != nil {
		log.Error().Err(err).Msg("Failed to assign spy cat to mission")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Cat assigned to a mission")
}

type AddTargetToMissionReq struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
