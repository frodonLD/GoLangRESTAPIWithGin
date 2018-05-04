package router

import (
	"github.com/frodonLD/GoLangRESTAPIWithGin/handler"

	gin "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	//router := gin.Default()

	// Set Release mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router without any middleware by default
	r := gin.New()
	r.HandleMethodNotAllowed = true

	// Middlewares set
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// https://github.com/gin-gonic/gin#using-middleware
	r.Use(gin.Recovery())

	// Simple group: /api/v1
	v1 := r.Group("/")
	{
		v1.GET("/_health", handler.HealthCheck)
		v1.GET("/logs", handler.GetAllNotificationsHandler)
		v1.GET("/logs/:id", handler.GetNotificationHandler)
	}
	return r
}

// Start launch the gin router
func Start() {
	router := setupRouter()
	router.Run(":8080")
}
