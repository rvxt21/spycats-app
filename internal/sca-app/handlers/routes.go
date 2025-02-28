package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rvxt21/sca-agency/internal/sca-app/middlewares"
)

func (h *Handler) RegisterRoutes(router *gin.Engine) {

	spyCatGroup := router.Group("/spycats")
	{
		spyCatGroup.POST("/", h.CreateSpyCat)
		spyCatGroup.GET("/", h.GetAllSpyCats)
		spyCatGroup.DELETE("/:id", middlewares.IdMiddleware(), h.DeleteSpyCat)
		spyCatGroup.PUT("/:id", middlewares.IdMiddleware(), h.UpdateSpyCat)
		spyCatGroup.GET("/:id", middlewares.IdMiddleware(), h.GetSpyCatById)
	}

}
