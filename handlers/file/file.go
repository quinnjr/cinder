package file

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/quinnjr/cinder"
)

// Handler implementation.
type Handler struct {
	Filepath  string
	Format    interface{}
	Timestamp interface{}
	mu        sync.Mutex
}

// DefaultLogFormat is the default logging format
var DefaultLogFormat = "[%s] [%s] %s"

// DefaultTimestampFormat is the default timestamp format
var DefaultTimestampFormat = "02 Jan 2006 03:04:05 MST"

// New returns a new Handler instance.
// Parameters format and timestamp must be strings or nil. Otherwiser, the Handler will panic.
func New(fp string, format interface{}, timestamp interface{}) *Handler {
	if format == nil {
		format = DefaultLogFormat
	}
	if timestamp == nil {
		timestamp = DefaultTimestampFormat
	}
	if reflect.TypeOf(format).Kind() != reflect.String {
		panic(errors.New("format must be a string"))
	}
	if reflect.TypeOf(timestamp).Kind() != reflect.String {
		panic(errors.New("timestamp must be a string"))
	}
	return &Handler{
		Filepath:  fp,
		Format:    format,
		Timestamp: timestamp,
	}
}

// HandleLog implements the Handler interface.
func (h *Handler) HandleLog(e *cinder.Entry) error {
	names := e.Fields.Keys()
	level := e.Level.String()
	timestamp := e.Timestamp

	fd, err := os.OpenFile(h.Filepath, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	h.mu.Lock()
	defer h.mu.Unlock()

	// New file contents
	fc := fmt.Sprintf(h.Format.(string), timestamp.Format(h.Timestamp.(string)), level, e.Message)

	for _, name := range names {
		if name == "source" {
			continue
		}
		fc = fc + fmt.Sprintf("%s=%v", name, e.Fields.Get(name))
	}

	// Add a newline to the end of the new file contents.
	fc = fmt.Sprintln(fc)

	fd.WriteString(fc)

	return nil
}
