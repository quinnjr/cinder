package json

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/quinnjr/cinder"
)

// Handler implementation.
type Handler struct {
	*json.Encoder
	mu sync.Mutex
}

// New returns a new Handler instance.
func New(w io.Writer) *Handler {
	return &Handler{
		Encoder: json.NewEncoder(w),
	}
}

// Default returns a new Handler instance with Stderr set
// as the writer output.
func Default() *Handler {
	return New(os.Stderr)
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.Encoder.Encode(e)
}
