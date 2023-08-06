package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type AnalyticsEventRequest struct {
	EventName string `json:"eventName"`
}

func (h *Handler) AnalyticsEvent(c *gin.Context) {
	userAgent := c.Request.Header.Get("User-Agent")

	logger := zerolog.Ctx(c.Request.Context()).With().Str("user_agent", userAgent).Logger()

	var req AnalyticsEventRequest
	if err := c.BindJSON(&req); err != nil {
		logger.Err(err).Msg("bind analytics request error")
	}

	logger.Info().Msgf("event: %s", req.EventName)

	c.Status(http.StatusNoContent)
}
