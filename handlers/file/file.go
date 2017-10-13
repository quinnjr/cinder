package file

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers"
)

// Handler implementation.
type Handler struct {
	File      io.Writer
	padding   uint
	format    string
	timestamp string
	mu        sync.Mutex
}

// New returns a new Handler instance.
// Parameters format and timestamp must be strings or nil. Otherwiser, the Handler will panic.
func New(f io.Writer) *Handler {
	return &Handler{
		File:      f,
		padding:   3,
		format:    handlers.DefaultFormat,
		timestamp: handlers.DefaultTimestamp,
	}
}

// SetFormat sets the log format for the file handler.
func (h *Handler) SetFormat(f string) { h.format = f }

// GetFormat returns the log format for the file handler.
func (h *Handler) GetFormat() string { return h.format }

// SetTimestamp sets the timestamp format for the file handler.
func (h *Handler) SetTimestamp(f string) { h.timestamp = f }

// GetTimestamp returns the timestamp format for the file handler.
func (h *Handler) GetTimestamp() string { return h.timestamp }

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) (err error) {
	names := e.Fields.Keys()
	level := strings.ToUpper(e.Level.String())
	timestamp := e.Timestamp

	h.mu.Lock()
	defer h.mu.Unlock()

	// New file contents
	fmt.Fprintf(h.File, h.format, level, timestamp.Format(h.timestamp), e.Message)

	for _, name := range names {
		if name == "source" {
			continue
		}
		fmt.Fprintf(h.File, "%*s%s=%v", h.padding, "", name, e.Fields[name])
	}

	// Add a newline to the end of the new file contents.
	fmt.Fprintln(h.File)

	return nil
}
