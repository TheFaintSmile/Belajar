package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}

func LogRequest(c *gin.Context) {
	// Log the entire request details
	request := c.Request
	log.Printf("Method: %s, URL: %s, Headers: %v", request.Method, request.URL, request.Header)
	// Log the request body, if any
    body, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        log.Println("Error reading request body:", err)
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }
    log.Println("Request Body:", string(body))
    c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
    c.Next()
}
