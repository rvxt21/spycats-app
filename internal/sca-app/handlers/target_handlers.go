package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"github.com/rvxt21/sca-agency/internal/sca-app/service"
)

type TargetsHandler struct {
	s service.TargetsService
}

func NewTargetsHandler(svc service.TargetsService) *TargetsHandler {
	return &TargetsHandler{
		s: svc,
	}
}

func (h *TargetsHandler) AddTargetToMission(c *gin.Context) {
	id := c.Value("id").(uint)

	var req AddTargetToMissionReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	target := models.Target{
		MissionID:   id,
		Name:        req.Name,
		Country:     req.Country,
		IsCompleted: false,
	}

	if err := h.s.AddTargetToMission(id, &target); err != nil {
		return
	}
}

type DeleteTargetReq struct {
	TargetId uint `json:"target_id"`
}

func (h *TargetsHandler) DeleteTarget(c *gin.Context) {
	mission_id := c.Value("id").(uint)

	var req DeleteTargetReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.s.DeleteTarget(mission_id, req.TargetId); err != nil {
		log.Error().Err(err).Msg("Failed to delete target")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, "Mission deleted")
}

type UpdateNotesReq struct {
	TargetId uint `json:"target_id"`
	Notes    string
}

func (h *TargetsHandler) UpdateNotes(c *gin.Context) {
	mission_id := c.Value("id").(uint)

	var req UpdateNotesReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.s.UpdateNotes(mission_id, req.TargetId, req.Notes); err != nil {
		log.Error().Err(err).Msg("Failed to update target")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, "Notes updated successfully")
}

type UpdateTargetStatusReqBody struct {
	TargetId    uint `json:"target_id"`
	IsCompleted bool `json:"is_completed"`
}

func (h *TargetsHandler) UpdateTargerStatus(c *gin.Context) {
	mission_id := c.Value("id").(uint)

	var req UpdateTargetStatusReqBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.s.UpdateTargerStatus(mission_id, req.TargetId, req.IsCompleted); err != nil {
		log.Error().Err(err).Msg("Failed to update target")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, "Status updated successfully")
}
