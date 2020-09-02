package demo

import (
	"errors"
	"fmt"
	"time"
	"use-gin/model/influxdb"

	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/models"
)

const cSQLPartEnd = "fill(0) tz('Asia/Shanghai')"

// HandleInfluxdbDemo a demo of handle influxdb
func HandleInfluxdbDemo(c *gin.Context) {

	// your specific logic

}

func writeInfluxdb() error {
	bp, err := influxdb.NewBatchPoints()
	if err != nil {
		return err
	}

	tags := map[string]string{}
	fields := map[string]interface{}{}
	time := time.Now()

	pt, err := influxdb.NewPoint(tags, fields, time)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	if err := influxdb.Write(bp); err != nil {
		return err
	}
	return nil
}

func readInfluxdb() (*models.Row, error) {
	sql := fmt.Sprintf("select * from ... %s", cSQLPartEnd)

	res, err := influxdb.Query(sql)
	if err != nil {
		return nil, err
	}

	if res[0].Series == nil {
		err := errors.New("nodata")
		return nil, err
	}
	return &res[0].Series[0], nil
}
