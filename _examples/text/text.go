package main

import (
	"errors"
	"os"
	"time"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/text"
)

func main() {

	// Instantiate a new logger
	logger := cinder.New(cinder.FatalLevel, text.Default())

	// Log an info message
	logger.Info("This is an info message")

	// Log a debug message (no output)
	logger.Debug("This debug message should not be output")

	// Log a warning message with a field
	logger.WithField("a thing", "a thing's value").Warn("This is a warning message")

	// Log an error message with a non-consequential error
	err := errors.New("This is an error's value")
	logger.WithError(err).Error("This is an error message")

	time.Sleep(1 * time.Second)

	logger.Info("This is a delayed message")

	os.Exit(0)
}
