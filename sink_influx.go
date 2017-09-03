package telemetry

import (
	"time"

	logrus "github.com/sirupsen/logrus"

	client "github.com/influxdata/influxdb/client/v2"
)

type InfluxSink struct {
	influxdb client.Client
	database string
}

func NewInfluxSink(addr, username, password, database string) (Sink, error) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	t := &InfluxSink{
		influxdb: c,
		database: database,
	}

	return t, nil
}

func (t *InfluxSink) Publish(event string, fields map[string]interface{}) error {
	// create a batch
	cfg := client.BatchPointsConfig{
		Precision: "s",
		Database:  t.database,
	}
	bp, err := client.NewBatchPoints(cfg)
	if err != nil {
		logrus.WithError(err).Warnf("Could not create batch")
		return err
	}

	// create new point
	pt, err := client.NewPoint(event, nil, fields, time.Now())
	if err != nil {
		logrus.WithError(err).Warnf("Could not create new point")
		return err
	}

	// add point to batch
	bp.AddPoint(pt)

	// write the batch
	if err := t.influxdb.Write(bp); err != nil {
		logrus.WithError(err).Warnf("Could not write batch")
		return err
	}

	return nil
}
