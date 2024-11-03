package logging

import (
	"github.com/google/uuid"
	"log/slog"
	"strings"
	"sync"
	"time"
)
import "github.com/gin-gonic/gin"

const cidKey = "cid"

var cidStorage sync.Map

func SetCID(cid string) {
	cidStorage.Store(cidKey, cid)
}

func GetCID() string {
	if cid, ok := cidStorage.Load(cidKey); ok {
		return cid.(string)
	}
	return ""
}

func ClearCID() {
	cidStorage.Delete(cidKey)
}

func LogRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		userAgent := c.Request.UserAgent()
		cid := c.Request.Header.Get("X-Correlation-ID")

		if cid == "" {
			cid = uuid.New().String() // Generate a new CID if not provided
			cid = strings.Replace(cid, "-", "", -1)
		}
		SetCID(cid)

		Info("API Called",
			slog.String("client_ip", clientIP),
			slog.String("method", method),
			slog.String("path", path),
			slog.String("user_agent", userAgent),
		)

		start := time.Now()
		c.Next()
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		Info("API Execution Completed",
			slog.Int("status", statusCode),
			slog.String("latency", latency.String()),
			slog.String("error", errorMessage),
		)
		ClearCID()
	}
}
