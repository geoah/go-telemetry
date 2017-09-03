package telemetry

import (
	"github.com/sirupsen/logrus"
)

func NewStdSink() (Sink, error) {
	t := &StdSink{}
	return t, nil
}

type StdSink struct{}

func (t *StdSink) Publish(event string, fields map[string]interface{}) error {
	logrus.WithFields(logrus.Fields(fields)).Infof("Telemetry event: %s", event)
	return nil
}
