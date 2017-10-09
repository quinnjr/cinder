package file

import (
	"fmt"
	"io"
	"sync"

	"github.com/quinnjr/cinder"
)

// Handler implementation.
type Handler struct {
	File            io.Writer
	Padding         uint
	Format          string
	TimestampFormat string
	mu              sync.Mutex
}

// New returns a new Handler instance.
// Parameters format and timestamp must be strings or nil. Otherwiser, the Handler will panic.
func New(f io.Writer) *Handler {
	return &Handler{
		File:            f,
		Padding:         3,
		Format:          "[%s] [%s] %s",
		TimestampFormat: "02 Jan 2006 03:04:05 MST",
	}
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) (err error) {
	names := e.Fields.Keys()
	level := e.Level.String()
	timestamp := e.Timestamp

	h.mu.Lock()
	defer h.mu.Unlock()

	// New file contents
	fmt.Fprintf(h.File, h.Format, timestamp.Format(h.TimestampFormat), level, e.Message)

	for _, name := range names {
		if name == "source" {
			continue
		}
		fmt.Fprintf(h.File, "%*s%s=%v", h.Padding, "", name, e.Fields[name])
	}

	// Add a newline to the end of the new file contents.
	fmt.Fprintln(h.File)

	return nil
}
