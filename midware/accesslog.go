package midware

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"use-gin/handler"
	"use-gin/logger"
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
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyLogWriter

		startTime := time.Now()

		c.Next()

		responseBody := bodyLogWriter.body.String()

		var responseCode string
		var responseMsg string
		var responseData interface{}

		if responseBody != "" {
			res := &handler.Response{}
			err := json.Unmarshal([]byte(responseBody), &res)
			if err == nil {
				responseCode = res.Code
				responseMsg = res.Message
				responseData = res.Data
			}
		}

		latencyTime := time.Since(startTime).Seconds()

		if c.Request.Method == "POST" {
			_ = c.Request.ParseForm()
		}

		requestID, ok := c.Get("X-Request-Id")
		if !ok {
			requestID = ""
		}

		requestURI := c.Request.URL.Path
		if c.Request.URL.RawQuery != "" {
			requestURI = c.Request.URL.Path + "?" + c.Request.URL.RawQuery
		}

		logger.Log.WithFields(logrus.Fields{
			"client_ip":       c.ClientIP(),
			"request_method":  c.Request.Method,
			"request_uri":     requestURI,
			"http_status":     c.Writer.Status(),
			"latency_time":    latencyTime,
			"request_proto":   c.Request.Proto,
			"request_referer": c.Request.Referer(),
			"request_body":    c.Request.PostForm.Encode(),
			"request_id":      requestID,
			"response_code":   responseCode,
			"response_msg":    responseMsg,
			"response_data":   responseData,
			"reqponse_ua":     c.Request.UserAgent(),
		}).Info("accesslog")
	}
}
