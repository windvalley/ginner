package influxdb

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"use-gin/config"
	"use-gin/logger"
)

var Client client.Client

func Init() {
	var err error
	Client, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     config.Config().Influxdb.Address,
		Username: config.Config().Influxdb.Username,
		Password: config.Config().Influxdb.DBName,
	})
	if err != nil {
		logger.Log.Errorf("connect influxdb error: %v", err)
	}
}

func Close() {
	Client.Close()
}

func NewPoint(
	tags map[string]string,
	fields map[string]interface{},
	time time.Time,
) (*client.Point, error) {
	pt, err := client.NewPoint(
		config.Config().Influxdb.Measurement,
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

func NewBatchPoints() (client.BatchPoints, error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Config().Influxdb.DBName,
		Precision: "s",
	})
	if err != nil {
		logger.Log.Errorf("influxdb client NewBatchPoints error: %v", err)
		return nil, err
	}

	return bp, nil
}

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
	if err := Client.Write(bp); err != nil {
		return err
	}
	return nil
}

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
		Database: config.Config().Influxdb.DBName,
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
