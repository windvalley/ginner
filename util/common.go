package util

import (
	"ginner/config"
	"reflect"

	"github.com/gin-gonic/gin"
)

// GetRequestID get X-Request-Id
func GetRequestID(c *gin.Context) string {
	requestID, ok := c.Get("X-Request-Id")
	if !ok {
		requestID = ""
	}

	return requestID.(string)
}

// GetRequestURI get request uri that include queries
func GetRequestURI(c *gin.Context) string {
	requestURI := c.Request.URL.Path
	if c.Request.URL.RawQuery != "" {
		requestURI = c.Request.URL.Path + "?" + c.Request.URL.RawQuery
	}

	return requestURI
}

// GetUsername get username
func GetUsername(c *gin.Context) string {
	username, ok := c.Get("key")
	if !ok {
		username = config.UsernameGuest
	}

	return username.(string)
}

// HasEntry determine whether an entry exists in a container(slice/array/map)
func HasEntry(entries interface{}, entry interface{}) bool {
	containerValue := reflect.ValueOf(entries)

	switch reflect.TypeOf(entries).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < containerValue.Len(); i++ {
			if containerValue.Index(i).Interface() == entry {
				return true
			}
		}
	case reflect.Map:
		if containerValue.MapIndex(reflect.ValueOf(entry)).IsValid() {
			return true
		}
	}

	return false
}
