package midware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"ginner/api"
	"ginner/config"
	"ginner/logger"
	"ginner/model"
	"ginner/util"
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

		requestBody := ""
		if c.Request.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			requestBody = string(bodyBytes)
		}

		c.Next()

		responseBody := bodyLogWriter.body.String()
		responseCode, responseMsg, responseData := parseResponseBody(responseBody)

		latencyTime := time.Since(startTime).Seconds()

		requestID := util.GetRequestID(c)
		requestURI := util.GetRequestURI(c)
		username := util.GetUsername(c)

		clientIP := c.ClientIP()
		httpStatusCode := c.Writer.Status()
		requestReferer := c.Request.Referer()
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
			"request_body":    requestBody,
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
			ReqBody:    requestBody,
			ReqReferer: requestReferer,
			UserAgent:  requestUA,
			ReqTime:    startTime,
			ReqLatency: latencyTime,
			HTTPStatus: httpStatusCode,
			ResCode:    responseCode,
			ResMessage: responseMsg,
			ResData:    string(resData),
		}

		go func() {
			if err := userOperationLog.Create(); err != nil {
				logger.Log.Errorf(
					"request id %s: user operation log create failed: %s",
					requestURI, err)
			}
		}()
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
