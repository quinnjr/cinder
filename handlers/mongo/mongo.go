package mongo

import (
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/quinnjr/cinder"
)

// Handler implementation.
type Handler struct {
	mu         sync.Mutex
	Collection *mgo.Collection
}

// entry is a conversion of the cinder.Entry stuct with bson fields set for reflection.
type entry struct {
	Fields    cinder.Fields `bson:"fields,omitempty"`
	Level     string        `bson:"level"`
	Message   string        `bson:"message"`
	Timestamp time.Time     `bson:"timestamp"`
}

// marshal returns a properly formated BSON representation
// of the log entry.
func newEntry(e *cinder.Entry) *entry {
	return &entry{
		Fields:    e.Fields,
		Level:     e.Level.String(),
		Message:   e.Message,
		Timestamp: e.Timestamp,
	}
}

// New creates a new Handler instance.
func New(collection *mgo.Collection) *Handler {
	return &Handler{Collection: collection}
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	m := newEntry(e)

	return h.Collection.Insert(m)
}
