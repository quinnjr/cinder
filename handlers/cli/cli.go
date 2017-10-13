package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/quinnjr/cinder"
)

const DefaultTimestamp = "02-Jan-06 03:04:05 MST"

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

// Default ...
func Default() *Handler {
	return &Handler{
		Writer:  os.Stderr,
		Padding: 3,
	}
}

// HandleLog ...
func (h *Handler) HandleLog(e *cinder.Entry) error {
	color := colors[e.Level]
	level := strings.ToUpper(e.Level.String())
	names := e.Fields.Keys()

	h.mu.Lock()
	defer h.mu.Unlock()

	fmt.Fprintf(h.Writer, "%s \033[%dm%*s\033[0m %s%*s", e.Timestamp.Format(DefaultTimestamp), color, h.Padding+1, level, e.Message, h.Padding+1, "")

	for k, name := range names {
		if name == "source" {
			continue
		}
		if k == 0 {
			fmt.Fprintf(h.Writer, "\033[%dm%s\033[0m=%v", color, name, e.Fields[name])
		} else {
			fmt.Fprintf(h.Writer, "%*s\033[%dm%s\033[0m=%v%*s", h.Padding+1, "", color, name, e.Fields[name], h.Padding+1, "")
		}
	}

	fmt.Fprintln(h.Writer)

	return nil

}
