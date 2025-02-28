package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rvxt21/sca-agency/internal/sca-app/middlewares"
)

func (h *CatHandler) RegisterRoutes(router *gin.Engine) {

	spyCatGroup := router.Group("/spycats")
	{
		spyCatGroup.POST("/", h.CreateSpyCat)
		spyCatGroup.GET("/", h.GetAllSpyCats)
		spyCatGroup.DELETE("/:id", middlewares.IdMiddleware(), h.DeleteSpyCat)
		spyCatGroup.PATCH("/:id", middlewares.IdMiddleware(), h.UpdateSpyCat)
		spyCatGroup.GET("/:id", middlewares.IdMiddleware(), h.GetSpyCatById)
	}

}

func (h *MissionHandler) RegisterRoutesM(router *gin.Engine) {
	missionGroup := router.Group("/missions")
	missionGroup.POST("/", h.CreateMission)
	missionGroup.GET("/", h.GetAllMissions)

	missionGroup.DELETE("/:id", middlewares.IdMiddleware(), h.DeleteMission)
	missionGroup.PATCH("/:id/updatestatus", middlewares.IdMiddleware(), h.UpdateMissionStatus)
	missionGroup.PATCH("/:id/assigncat", middlewares.IdMiddleware(), h.SetCatForMission)
	missionGroup.GET("/:id", middlewares.IdMiddleware(), h.GetMission)

}

func (h *TargetsHandler) RegisterRoutesT(router *gin.Engine) {
	targetGroup := router.Group("/targets")
	targetGroup.POST("/", middlewares.IdMiddleware(), h.AddTargetToMission)
	targetGroup.DELETE("/", middlewares.IdMiddleware(), h.DeleteTarget)
}
