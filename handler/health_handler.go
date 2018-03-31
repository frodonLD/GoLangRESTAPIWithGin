package handler

import (
	"net/http"

	"github.com/frodonLD/GoLangRESTAPIWithGin/utils"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	lc := utils.GetLoggingContextFromGinContext(c)
	logger := lc.LoggerEntry
	logger.Warn("HealthCheckHandler called.")
	s := `a mutiline string
	line 2
	line3
	line 4
	line 5
	line 6
	`
	logger.Debug("Multiline log", s)
	c.JSON(http.StatusOK, gin.H{"alive": true})
}
