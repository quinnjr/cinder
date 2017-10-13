package text

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/quinnjr/cinder"
)

const (
	// DefaultTimestamp The default timestamp format
	DefaultTimestamp = "02-Jan-06 03:04:05 MST"
	// DefaultFormat The default output format
	DefaultFormat = "[%s] [%s] %s%*s"
)

// Handler implementation.
type Handler struct {
	Writer    io.Writer
	Padding   uint
	format    string
	timestamp string
	mu        sync.Mutex
}

// New returns a Handler instance.
func New(w io.Writer) *Handler {
	return &Handler{
		Writer:    w,
		Padding:   3,
		format:    DefaultFormat,
		timestamp: DefaultTimestamp,
	}
}

// Default returns a Handler instance using os.Stderr as the writer.
func Default() *Handler {
	return &Handler{Writer: os.Stderr, Padding: 3, format: DefaultFormat, timestamp: DefaultTimestamp}
}

// SetFormat sets the log format for the text handler.
func (h *Handler) SetFormat(f string) { h.format = f }

// GetFormat returns the log format for the text handler.
func (h *Handler) GetFormat() string { return h.format }

// SetTimestamp sets the timestamp format for the text handler.
func (h *Handler) SetTimestamp(f string) { h.timestamp = f }

// GetTimestamp returns the timestamp format for the text handler.
func (h *Handler) GetTimestamp() string { return h.timestamp }

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) (err error) {
	level := strings.ToUpper(e.Level.String())
	names := e.Fields.Keys()
	timestamp := e.Timestamp.Format(h.timestamp)

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err = fmt.Fprintf(h.Writer, h.format, level, timestamp, e.Message, h.Padding*3, "")
	if err != nil {
		return
	}

	for _, name := range names {
		_, err = fmt.Fprintf(h.Writer, "%*s%s=%v", h.Padding, "", name, e.Fields[name])
		if err != nil {
			return
		}
	}

	_, err = fmt.Fprintln(h.Writer)

	return
}
