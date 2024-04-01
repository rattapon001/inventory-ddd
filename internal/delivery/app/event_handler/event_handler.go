package eventhandler

type EventHandler interface {
	Publish(eventName string, payload interface{}) error
}
