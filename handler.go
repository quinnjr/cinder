package cinder

// HandlerFunc allows for the use of a custom function as a log handler.
type HandlerFunc func(*Entry) error

// Handler is used to handle logging events.
// Each Handler must implement the Handler interface and is
// responsible for implementing the logic of handling events.
type Handler interface {
	HandleLog(*Entry) error
}

// HandleLog calls the handler's HandleLog function.
func (hf HandlerFunc) HandleLog(e *Entry) error {
	return hf(e)
}
