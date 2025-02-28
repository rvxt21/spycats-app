package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)


func IdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id") 
		if idStr == "" {
			log.Info().Msg("Missed ID")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			log.Warn().Err(err).Msg("Invalid ID param")
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Set("id", uint(id))

		c.Next()
	}
}