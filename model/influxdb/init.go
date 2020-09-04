package influxdb

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"use-gin/config"
	"use-gin/logger"
)

// Client client instance of influxdb
var Client client.Client

// Init influxdb initialization
func Init() {
	var err error
	Client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Conf().Influxdb.Address,
		Username: config.Conf().Influxdb.Username,
		Password: config.Conf().Influxdb.DBName,
	})
	if err != nil {
		logger.Log.Errorf("connect influxdb error: %v", err)
	}
}

// Close close connections of influxdb
func Close() {
	Client.Close()
}

// NewPoint get *client.Point
func NewPoint(
	tags map[string]string,
	fields map[string]interface{},
	time time.Time,
) (*client.Point, error) {
	pt, err := client.NewPoint(
		config.Conf().Influxdb.Measurement,
		tags,
		fields,
		time,
	)
	if err != nil {
		logger.Log.Errorf("influxdb client NewPoint error: %v", err)
		return nil, err
	}

	return pt, nil
}

// NewBatchPoints get client.BatchPoints
func NewBatchPoints() (client.BatchPoints, error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Conf().Influxdb.DBName,
		Precision: "s",
	})
	if err != nil {
		logger.Log.Errorf("influxdb client NewBatchPoints error: %v", err)
		return nil, err
	}

	return bp, nil
}

// Write write data into influxdb
//    e.g.
// bp, err := influxdb.NewBatchPoints()
//if err != nil {
//return err
//}
//tags := map[string]string{}
//fields := map[string]interface{}{}
//time := time.Now()
//pt, err := influxdb.NewPoint(tags, fields, time)
//if err != nil {
//return err
//}
//bp.AddPoint(pt)
//if err := influxdb.Write(bp); err != nil {
//return err
//}
//return nil
func Write(bp client.BatchPoints) error {
	err := Client.Write(bp)
	return err
}

// Query get data from influxdb
//     e.g.
//res, err := influxdb.Query(sql)
//if err != nil {
//return nil, err
//}
//if res[0].Series == nil {
//err := errors.New("nodata")
//return nil, err
//}
//return res[0].Series[0], nil
func Query(cmd string) ([]client.Result, error) {
	q := client.Query{
		Command:  cmd,
		Database: config.Conf().Influxdb.DBName,
	}

	response, err := Client.Query(q)
	if err != nil {
		return nil, err
	}

	if response.Error() != nil {
		return nil, response.Error()
	}

	return response.Results, nil
}
