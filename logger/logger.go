package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLogger() gin.HandlerFunc {
	const format = "%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n"

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(format,
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
