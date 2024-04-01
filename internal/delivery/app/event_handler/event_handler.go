package eventhandler

type EventHandler interface {
	publish(eventName string, payload interface{}) error
}
