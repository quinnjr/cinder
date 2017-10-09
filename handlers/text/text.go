package text

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/quinnjr/cinder"
)

var start = time.Now()

// Handler implementation.
type Handler struct {
	mu     sync.Mutex
	Writer io.Writer
}

// New returns a Handler instance.
func New(w io.Writer) *Handler {
	return &Handler{Writer: w}
}

// Default returns a Handler instance using os.Stderr as the writer.
func Default() *Handler {
	return &Handler{Writer: os.Stderr}
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) error {
	level := strings.ToUpper(e.Level.String())
	names := e.Fields.Keys()

	h.mu.Lock()
	defer h.mu.Unlock()

	since := time.Since(start) / time.Second

	fmt.Fprintf(h.Writer, "[%s] [%04d] %-20s", level, since, e.Message)

	for _, name := range names {
		fmt.Fprintf(h.Writer, "%s=%-3v", name, e.Fields[name])
	}

	fmt.Fprintln(h.Writer)

	return nil
}
