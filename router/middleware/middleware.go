package middleware

import (
	"github.com/frodonLD/GoLangRESTAPIWithGin/utils"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// AddCorrelationID check Correlation ID existance in the Header of the request. Set it if non-exist
// https://github.com/gin-gonic/gin#custom-middleware
func AddCorrelationID() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationID := c.Request.Header.Get(utils.CorrelationIDKey)
		if correlationID == "" {
			correlationID = uuid.NewV4().String()
		}
		c.Set(utils.CorrelationIDKey, correlationID)
		c.Next()
	}
}

// AddLoggingContext set a logger based on the gin context
// https://github.com/gin-gonic/gin#custom-middleware
func AddLoggingContext(lc *utils.LoggingContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get correlationId from the context
		correlationID, ok := c.Get(utils.CorrelationIDKey)
		if !ok {
			correlationID = ""
		}
		lc.SetCorrelationID(correlationID.(string))
		c.Set(utils.LoggingContextKey, lc)
		c.Next()
	}
}
