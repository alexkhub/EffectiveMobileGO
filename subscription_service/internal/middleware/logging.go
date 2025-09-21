package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func Logging() gin.HandlerFunc{
	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		req_id := fmt.Sprintf("req-%s", uuid.New().String())

		c.Set("req_id", req_id)

		c.Next()
		latency := time.Since(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		
		message, ok := c.Get("message")
		if !ok{
			message = ""
		}
		if raw != "" {
			path = path + "?" + raw
		}
		if statusCode >= 100 && statusCode < 400{
			logrus.WithFields(logrus.Fields{
				"method": method,
				"path": path,
				"status": statusCode,
				"request_id": req_id,
				"clientIP": clientIP,
				"start": start.Format(time.DateTime),
				"latency": latency,
			}).Info(message.(string))

		}else if statusCode >= 400 && statusCode < 500{
			logrus.WithFields(logrus.Fields{
				"method": method,
				"path": path,
				"status": statusCode,
				"request_id": req_id,
				"clientIP": clientIP,
				"start": start.Format(time.DateTime),
				"latency": latency,
			}).Warn(message.(string))
			
		}else{
			logrus.WithFields(logrus.Fields{
				"method": method,
				"path": path,
				"status": statusCode,
				"request_id": req_id,
				"clientIP": clientIP,
				"start": start.Format(time.DateTime),
				"latency": latency,
			}).Error(message.(string))
		}
	}
}
