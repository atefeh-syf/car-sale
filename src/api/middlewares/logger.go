package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/atefeh-syf/car-sale/config"
	"github.com/atefeh-syf/car-sale/pkg/logging"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.ResponseWriter.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc{
	logger := logging.NewLogger(cfg)
	return StructuredLogger(logger)
}

func StructuredLogger(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	
		start := time.Now()
		path := c.FullPath()
		raw := c.Request.URL.RawQuery

		bodyByte, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
		
		c.Writer = blw
		c.Next()

		params := gin.LogFormatterParams{}

		params.TimeStamp = time.Now()
		params.Latency = time.Now().Sub(start)
		params.ClientIP = c.ClientIP()
		params.Method = c.Request.Method
		params.StatusCode = c.Writer.Status()
		params.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		params.BodySize = c.Writer.Size()


		if raw != "" {
			path  = path + "?" + raw
		}
		params.Path = path
		
		keys := map[logging.ExtraKey]interface{}{}
		keys[logging.Path] = params.Path
		keys[logging.TimeStamp] = params.TimeStamp
		keys[logging.Latency] = params.Latency
		keys[logging.ClientIp] = params.ClientIP
		keys[logging.Method] = params.Method
		keys[logging.StatusCode] = params.StatusCode
		keys[logging.ErrorMessage] = params.ErrorMessage
		keys[logging.BodySize] = params.BodySize
		keys[logging.RequestBody] = string(bodyByte)
		keys[logging.ResponseBody] = blw.body.String()

		logger.Info(logging.RequestResponse, logging.Api, "", keys)
	}	
}