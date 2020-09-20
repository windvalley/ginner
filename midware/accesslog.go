package midware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"use-gin/api"
	"use-gin/config"
	"use-gin/logger"
	"use-gin/model"
	"use-gin/util"
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
		responseCode, responseMsg, responseData := parseResponseBody(responseBody)

		latencyTime := time.Since(startTime).Seconds()

		if c.Request.Method == http.MethodPost {
			_ = c.Request.ParseForm()
		}

		requestID := util.GetRequestID(c)
		requestURI := util.GetRequestURI(c)
		username := util.GetUsername(c)

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

		if !doUserAudit(c, username) {
			return
		}

		resData, err := json.Marshal(responseData)
		if err != nil {
			resData = []byte("")
		}

		userOperationLog := &model.UserOperationLog{
			Username:   username,
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

func parseResponseBody(responseBody string) (
	responseCode, responseMsg string, responseData interface{}) {
	if responseBody != "" {
		res := &api.Response{}
		err := json.Unmarshal([]byte(responseBody), &res)
		if err != nil {
			return "", "", nil
		}

		responseCode = res.Code
		responseMsg = res.Message
		responseData = res.Data
		return
	}

	return "", "", nil
}

func doUserAudit(c *gin.Context, username string) bool {
	isUserAudit, ok := c.Get(config.UserAuditEnableKey)
	if !ok || isUserAudit != true {
		return false
	}

	if username == config.UsernameGuest || c.Request.Method == http.MethodGet ||
		c.Request.Method == http.MethodOptions {
		return false
	}

	return true
}
