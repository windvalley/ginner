package demo1

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/models"

	"use-gin/model/influxdb"
)

const SQL_PART_END = "fill(0) tz('Asia/Shanghai')"

var influx = influxdb.New()

func HandleInfluxdbDemo(c *gin.Context) {

	// your specific logic

}

func WriteInfluxdb() error {
	client, err := influx.Connect()
	if err != nil {
		return err
	}
	defer client.Close()

	bp, err := influx.NewBatchPoints()
	if err != nil {
		return err
	}

	tags := map[string]string{}
	fields := map[string]interface{}{}
	time := time.Now()

	pt, err := influx.NewPoint(tags, fields, time)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	if err := client.Write(bp); err != nil {
		return err
	}

	return nil
}

func ReadInfluxdb() (*models.Row, error) {
	client, err := influx.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	sql := fmt.Sprintf("select * from ... %s", SQL_PART_END)

	res, err := influx.QueryDB(client, sql)
	if err != nil {
		return nil, err
	}

	if res[0].Series == nil {
		err := errors.New("nodata")
		return nil, err
	}
	return &res[0].Series[0], nil
}
