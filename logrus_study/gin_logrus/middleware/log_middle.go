package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		end := time.Now()
		timesub := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		//bodySize := c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		logrus.Print("[GIN] %s | %d | %d | %s | %s |%s",
			start.Format("2006-01-02 12:34:02"),
			statusCode,
			timesub,
			clientIP,
			method,
			path,
		)
	}
}
