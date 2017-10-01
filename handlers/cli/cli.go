package cli

import (
	"fmt"
	"io"
	"sync"

	"github.com/quinnjr/cinder"
)

// Colors
const (
	none   = 0
	red    = 31
	green  = 32
	yellow = 33
	blue   = 34
	grey   = 37
)

var colors = []int{
	cinder.DebugLevel: grey,
	cinder.ErrorLevel: red,
	cinder.FatalLevel: red,
	cinder.InfoLevel:  blue,
	cinder.WarnLevel, yellow,
}

// Handler ...
type Handler struct {
	mu      sync.Mutex
	Writer  io.Writer
	Padding int
}

// New ...
func New(w io.Writer) *Handler {
	return &Handler{
		Writer:  w,
		Padding: 3,
	}
}

// HandleLog ...
func (h *Handler) HandleLog(e *cinder.Entry) error {
	color := colors[e.Level]
	level := e.Level.String()
	names := e.Fields.Keys()

	h.mu.Lock()
	defer h.mu.Unlock()

	fmt.Fprintf(h.Writer, "\033[%dm%*s\033[0m %-25s", color, h.Padding+1, level, e.Message)

	for _, name := range names {
		if name == "source" {
			continue
		}

		fmt.Fprintf(h.Writer, " \033[%dm%s\033[0m=%v", color, name, e.Fields.Get(name))
	}

	fmt.Fprintln(h.Writer)

	return nil

}
