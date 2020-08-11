package influxdb

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"use-gin/config"
	"use-gin/logger"
)

type InfluxDB struct {
	Address     string
	Username    string
	Password    string
	DBName      string
	Measurement string
}

func New() *InfluxDB {
	return &InfluxDB{
		Address:     config.Config().Influxdb.Address,
		Username:    config.Config().Influxdb.Username,
		Password:    config.Config().Influxdb.Password,
		DBName:      config.Config().Influxdb.DBName,
		Measurement: config.Config().Influxdb.Measurement,
	}
}

func (i *InfluxDB) Connect() (client.Client, error) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     i.Address,
		Username: i.Username,
		Password: i.Password,
	})
	if err != nil {
		logger.Log.Errorf("connect influxdb error: %v", err)
		return nil, err
	}

	return c, nil
}

func (i *InfluxDB) NewPoint(
	tags map[string]string,
	fields map[string]interface{},
	time time.Time,
) (*client.Point, error) {
	pt, err := client.NewPoint(i.Measurement, tags, fields, time)
	if err != nil {
		logger.Log.Errorf("influxdb client NewPoint error: %v", err)
		return nil, err
	}

	return pt, nil
}

func (i *InfluxDB) NewBatchPoints() (client.BatchPoints, error) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.DBName,
		Precision: "s",
	})
	if err != nil {
		logger.Log.Errorf("influxdb client NewBatchPoints error: %v", err)
		return nil, err
	}

	return bp, nil
}

//     e.g.
//res, err := i.QueryDB(c, sql)
//if err != nil {
//return nil, err
//}
//if res[0].Series == nil {
//err := errors.New("nodata")
//return nil, err
//}
//return res[0].Series[0], nil
func (i *InfluxDB) QueryDB(
	c client.Client,
	cmd string,
) ([]client.Result, error) {
	q := client.Query{
		Command:  cmd,
		Database: i.DBName,
	}

	response, err := c.Query(q)
	if err != nil {
		return nil, err
	}

	if response.Error() != nil {
		return nil, response.Error()
	}

	return response.Results, nil
}
