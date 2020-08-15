package demo1

import (
	"errors"
	"fmt"
	"time"
	"use-gin/model/influxdb"

	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/models"
)

const SQL_PART_END = "fill(0) tz('Asia/Shanghai')"

func HandleInfluxdbDemo(c *gin.Context) {

	// your specific logic

}

func WriteInfluxdb() error {
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

func ReadInfluxdb() (*models.Row, error) {
	sql := fmt.Sprintf("select * from ... %s", SQL_PART_END)

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
