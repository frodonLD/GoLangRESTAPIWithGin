package handler

import (
	"net/http"

	"github.com/frodonLD/GoLangRESTAPIWithGin/utils"
	"github.com/sirupsen/logrus"

	"github.com/frodonLD/GoLangRESTAPIWithGin/model"
	"github.com/gin-gonic/gin"
)

func GetAllNotificationsHandler(c *gin.Context) {
	// lc := utils.GetLoggingContextFromGinContext(c)
	// logger := lc.LoggerEntry
	// logger.Debug("GetLogsHandler called.")
	c.JSON(http.StatusOK, model.Notifications)
}

func GetNotificationHandler(c *gin.Context) {
	lc := utils.GetLoggingContextFromGinContext(c)
	logger := lc.LoggerEntry
	logger.WithFields(logrus.Fields{
		"params": c.Params,
	}).Debug("GetLogHandler called.")
	// https://github.com/gin-gonic/gin#parameters-in-path
	id := c.Param("id")
	for _, item := range model.Notifications {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
