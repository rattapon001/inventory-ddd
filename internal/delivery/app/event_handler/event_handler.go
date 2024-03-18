package eventhandler

type EventHandler interface {
	publish(event interface{}) error
}
