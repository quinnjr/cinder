package google

import (
	"fmt"
	"sync"

	"golang.org/x/net/context"

	"cloud.google.com/go/logging"
	"github.com/quinnjr/cinder"
)

// Handler implementation.
type Handler struct {
	Logger *logging.Logger
	mu     sync.Mutex
}

// New returns a new Handler instance.
func New(project string, logName string, options logging.LoggerOption) *Handler {
	ctx := context.Background()
	client, _ := logging.NewClient(ctx, project)
	return &Handler{Logger: client.Logger(logName)}
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	p := fmt.Sprintf("%-25s", e.Message)

	for k, v := range e.Fields {
		p = p + fmt.Sprintf("%s=%-3v", k, v)
	}

	h.Logger.Log(logging.Entry{
		Timestamp: e.Timestamp,
		Severity:  logging.Severity(e.Level),
		Payload:   p,
	})

	return nil
}
