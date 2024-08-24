package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

// CustomResponseWriter is a custom response writer that captures the response.
type CustomResponseWriter struct {
	gin.ResponseWriter
	ResponseBody *bytes.Buffer
	StatusCode   int
}

// Write is an implementation of the io.Writer interface that captures the response body.
func (r *CustomResponseWriter) Write(b []byte) (int, error) {
	// Capture the response body
	r.ResponseBody.Write(b)
	// Write to the original response writer
	return r.ResponseWriter.Write(b)
}

// WriteHeader is an implementation of the http.ResponseWriter interface that captures the status code.
func (r *CustomResponseWriter) WriteHeader(statusCode int) {
	// Capture the status code
	r.StatusCode = statusCode
	// Write to the original response writer
	r.ResponseWriter.WriteHeader(statusCode)
}

// CaptureResponseMiddleware is a Gin middleware that captures the response body.
func CaptureResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a ResponseCapturingWriter and replace the original writer
		c.Writer = &CustomResponseWriter{
			ResponseWriter: c.Writer,
			ResponseBody:   bytes.NewBufferString(""),
		}

		c.Next()
	}
}
