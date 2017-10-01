package memory

import (
	"sync"

	"github.com/quinnjr/cinder"
)

// Handler implementation.
type Handler struct {
	mu      sync.Mutex
	Entries []*cinder.Entry
}

// New creates a new Handler instance.
func New() *Handler {
	return &Handler{}
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Entries = append(h.Entries, e)
	return nil
}
