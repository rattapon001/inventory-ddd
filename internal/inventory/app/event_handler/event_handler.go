package eventhandler

// EventHandler interface
type EventHandler interface {
	publish(event interface{}) error
}
