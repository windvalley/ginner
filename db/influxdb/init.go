package influxdb

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"ginner/config"
	"ginner/logger"
)

var cli client.Client

// Init influxdb initialization
func Init() {
	var err error
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Conf().Influxdb.Address,
		Username: config.Conf().Influxdb.Username,
		Password: config.Conf().Influxdb.Password,
	})
	if err != nil {
		logger.Log.Fatalf("connect influxdb failed: %v", err)
	}
}

// Close close connections of influxdb
func Close() {
	if err := cli.Close(); err != nil {
		logger.Log.Errorf("close influxdb client error: %v", err)
	}
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

	return pt, err
}

// NewBatchPoints get client.BatchPoints
func NewBatchPoints() (client.BatchPoints, error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Conf().Influxdb.DBName,
		Precision: "s",
	})

	return bp, err
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
	err := cli.Write(bp)
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
func Query(sql string) ([]client.Result, error) {
	q := client.Query{
		Command:  sql,
		Database: config.Conf().Influxdb.DBName,
	}

	response, err := cli.Query(q)
	if err != nil {
		return nil, err
	}

	if response.Error() != nil {
		return nil, response.Error()
	}

	return response.Results, nil
}
