package telemetry

import "github.com/sirupsen/logrus"

func New() *Telemetry {
	sink, _ := NewStdSink()
	return &Telemetry{
		sink: sink,
	}
}

type Telemetry struct {
	sink Sink
}

func (t *Telemetry) SetSink(sink Sink) {
	t.sink = sink
}

func (t *Telemetry) Publish(event string, fields map[string]interface{}) error {
	logrus.WithField("event", event).Infof("Publishing telemetry")
	return t.sink.Publish(event, fields)
}
