package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom logger - we can customise as we want
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(gl gin.LogFormatterParams) string {
		clientIp := fmt.Sprintf("%s", gl.ClientIP)
		if clientIp == "::1" {
			clientIp = "121.0.0.1"
		}

		return fmt.Sprintf("ACCESS-LOG - [%s] | IP[%s] | %d | %s | %s | %s\n",
			gl.TimeStamp.Format(time.RFC3339),
			clientIp,
			gl.StatusCode,
			gl.Latency,
			gl.Method,
			gl.Path,
		)
	})
}
