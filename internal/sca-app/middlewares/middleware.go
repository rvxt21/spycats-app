package middlewares

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
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

func LogMiddleware() gin.HandlerFunc {

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return func(c *gin.Context) {

		start := time.Now()

		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Str("user_agent", c.Request.UserAgent()).
			Msg("Request received")
		c.Next()

		duration := time.Since(start)

		entry := logger.With().
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Int("status", c.Writer.Status()).
			Str("referrer", c.Request.Referer()).
			Str("request_id", c.Writer.Header().Get("Request-Id")).
			Dur("duration", duration).
			Logger()

		if c.Writer.Status() >= 500 {
			entry.Error().Str("error", c.Errors.String()).Msg("Request failed")
		} else {
			entry.Info().Msg("Request completed")
		}
	}
}
