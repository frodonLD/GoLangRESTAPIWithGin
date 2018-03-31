package handler

import (
	"net/http"

	"github.com/frodonLD/GoLangRESTAPIWithGin/model"
	"github.com/gin-gonic/gin"
)

func GetAllNotificationsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Notifications)
}

func GetNotificationHandler(c *gin.Context) {
	id := c.Param("id")
	for _, item := range model.Notifications {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
