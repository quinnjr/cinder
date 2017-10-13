package main

import (
	"errors"
	"os"
	"time"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/cli"
)

func main() {

	// Instantiate a new logger
	logger := cinder.New(cinder.FatalLevel, cli.Default())

	// Log an info message
	logger.Info("This is an info message")

	// Log a debug message (no output)
	logger.Debug("This debug message should not be output")

	// Log a warning message with a field
	logger.WithFields(cinder.Fields{
		"a thing":       "a thing's value",
		"another thing": false,
	}).Warn("This is a warning message")

	// Log an error message with a non-consequential error
	err := errors.New("This is an error's value")
	logger.WithError(err).Error("This is an error message")

	time.Sleep((1 * time.Second) + (2 * time.Millisecond))

	logger.Info("This is a delayed message")

	os.Exit(0)
}
