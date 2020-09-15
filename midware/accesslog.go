package midware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"use-gin/handler"
	"use-gin/logger"
	"use-gin/model/rdb"
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

// AccessLogger write access log to log file and write user operation log to db.
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

		username, ok := c.Get("key")
		if !ok {
			username = "guest"
		}

		clientIP := c.ClientIP()
		httpStatusCode := c.Writer.Status()
		requestReferer := c.Request.Referer()
		requestBodyData := c.Request.PostForm.Encode()
		requestUA := c.Request.UserAgent()

		logger.Log.WithFields(logrus.Fields{
			"username":        username,
			"client_ip":       clientIP,
			"request_method":  c.Request.Method,
			"request_uri":     requestURI,
			"http_status":     httpStatusCode,
			"latency_time":    latencyTime,
			"request_proto":   c.Request.Proto,
			"request_referer": requestReferer,
			"request_body":    requestBodyData,
			"request_id":      requestID,
			"request_ua":      requestUA,
			"response_code":   responseCode,
			"response_msg":    responseMsg,
		}).Info("accesslog")

		if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodOptions {
			return
		}

		resData, err := json.Marshal(responseData)
		if err != nil {
			resData = []byte("")
		}

		userOperationLog := &rdb.UserOperationLog{
			Username:   username.(string),
			ClientIP:   clientIP,
			ReqMethod:  c.Request.Method,
			ReqPath:    requestURI,
			ReqBody:    requestBodyData,
			ReqReferer: requestReferer,
			UserAgent:  requestUA,
			ReqTime:    startTime,
			ReqLatency: latencyTime,
			HTTPStatus: httpStatusCode,
			ResCode:    responseCode,
			ResMessage: responseMsg,
			ResData:    string(resData),
		}

		go userOperationLog.Create()
	}
}
