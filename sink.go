package telemetry

// Sink interface for publishing events
type Sink interface {
	Publish(event string, fields map[string]interface{}) error
}
