package telemetry

var t *Telemetry

func init() {
	t = New()
}

func Publish(event string, fields map[string]interface{}) error {
	return t.Publish(event, fields)
}

func SetSink(sink Sink) {
	t.sink = sink
}
