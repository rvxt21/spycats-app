package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	catapi "github.com/rvxt21/sca-agency/external/cat_api"
	"github.com/rvxt21/sca-agency/internal/sca-app/models"
	"github.com/rvxt21/sca-agency/internal/sca-app/service"
)

type CatHandler struct {
	s        service.Service
	breedAPI *catapi.BreedAPIChecker
}

func New(s service.Service, b *catapi.BreedAPIChecker) *CatHandler {
	return &CatHandler{
		s:        s,
		breedAPI: b,
	}
}

func (h *CatHandler) CreateSpyCat(c *gin.Context) {
	var cat models.SpyCat

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if !h.breedAPI.CheckIfBreedExists(cat.Breed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Breed is not valid"})
		return
	}

	if err := h.s.CreateSpyCat(&cat); err != nil {
		log.Error().Err(err).Msg("Failed to create spy cat")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create spy cat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Spy cat created successfully"})
}

func (h *CatHandler) DeleteSpyCat(c *gin.Context) {
	id := c.Value("id").(uint)

	if err := h.s.DeleteSpyCat(id); err != nil {
		log.Error().Err(err).Msg("Failed to delete spy cat")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete spy cat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Spy cat deleted successfully"})
}

type UpdateSalaryReq struct {
	salary float64 `json:"salary"`
}

func (h *CatHandler) UpdateSpyCat(c *gin.Context) {
	var req UpdateSalaryReq

	id := c.Value("id").(uint)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.s.UpdateSalary(id, req.salary); err != nil {
		log.Error().Err(err).Msg("Failed to update salary")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update salary"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Salary updated successfully"})

}

func (h *CatHandler) GetAllSpyCats(c *gin.Context) {
	cats, err := h.s.GetAllSpyCats()
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve all spy cats")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve spy cats"})
		return
	}

	c.JSON(http.StatusOK, cats)
}

func (h *CatHandler) GetSpyCatById(c *gin.Context) {
	id := c.Value("id").(uint)

	cat, err := h.s.GetSpyCat(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve spy cat by ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve spy cat"})
		return
	}

	c.JSON(http.StatusOK, cat)
}
