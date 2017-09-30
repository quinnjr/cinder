package text

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/quinnjr/cinder"
)

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
	level := e.Level.String()
	names := e.Fields.Fields()
	start := time.Now()

	h.mu.Lock()
	defer h.mu.Unlock()

	_, err := fmt.Fprintf(h.Writer, "[%s][%d] %-25s", level, time.Since(start)/time.Second, e.Message)
	if err != nil {
		return err
	}

	for _, name := range names {
		_, err = fmt.Fprintf(h.Writer, "%s=%v", name, e.Fields.Get(name))
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintln(h.Writer)
	if err != nil {
		return err
	}
	return nil

}
