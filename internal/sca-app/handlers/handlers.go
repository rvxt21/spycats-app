package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"github.com/rvxt21/sca-agency/internal/sca-app/service"
)

type Handler struct {
	s service.Service
}

func (h *Handler) CreateSpyCat(c *gin.Context) {
	var cat models.SpyCat

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

}

func (h *Handler) DeleteSpyCat(c *gin.Context) {
	// id := c.Value("id").(int)

}

type UpdateSalaryReq struct {
	salary float64 `json:"salary"`
}

func (h *Handler) UpdateSpyCat(c *gin.Context) {
	var req UpdateSalaryReq

	// id := c.Value("id").(int)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

}

func (h *Handler) GetAllSpyCats(c *gin.Context) {

}

func (h *Handler) GetSpyCatById(c *gin.Context) {
	// id := c.Value("id").(int)

}
