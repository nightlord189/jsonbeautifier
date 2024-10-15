package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nightlord189/jsonbeautifier/internal/config"
)

type Handler struct {
	Config config.HTTPConfig
}

func New(cfg config.HTTPConfig) *Handler {
	return &Handler{Config: cfg}
}

func (h *Handler) Run() error {
	gin.SetMode(h.Config.GinMode)

	router := gin.New()
	router.Use(gin.Recovery())
	ginLoggerConfig := gin.LoggerConfig{}

	if h.Config.GinMode == "release" {
		ginLoggerConfig.Output = io.Discard
	}

	router.Use(gin.LoggerWithConfig(ginLoggerConfig))

	corsMiddleware := cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Access-Control-Allow-Origin", "Content-Type", "Content-Length",
			"Accept-Encoding", "Authorization", "X-CSRF-Token", "X-Request-ID", "X-Forwarded-For", "Origin", "Referer",
		},
		ExposeHeaders:    []string{"Content-Length", "Content-Range", "X-Request-ID", "X-Forwarded-For", "Origin", "Referer"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	})

	router.Use(corsMiddleware)

	router.OPTIONS("/", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	router.LoadHTMLFiles("web/template/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	return router.Run(fmt.Sprintf(":%d", h.Config.Port))
}
