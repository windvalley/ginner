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
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
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

		latencyTime := time.Since(startTime).Milliseconds()

		if c.Request.Method == "POST" {
			_ = c.Request.ParseForm()
		}

		logger.Log.WithFields(logrus.Fields{
			"client_ip":    c.ClientIP(),
			"req_method":   c.Request.Method,
			"req_uri":      c.Request.URL.Path,
			"http_status":  c.Writer.Status(),
			"latency_time": latencyTime,
			"req_proto":    c.Request.Proto,
			"req_referer":  c.Request.Referer(),
			"req_body":     c.Request.PostForm.Encode(),
			"res_code":     responseCode,
			"res_msg":      responseMsg,
			"res_data":     responseData,
			"req_ua":       c.Request.UserAgent(),
		}).Debug()
	}
}
