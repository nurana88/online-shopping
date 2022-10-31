package logger

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

type (
	IndexContext int
	// Logger defines a logging interface
	Logger logrus.FieldLogger
)

const (
	TraceIDKey IndexContext = iota
	MessageIDKey
)

var (
	baseLogger logrus.FieldLogger
)

func init() {
	baseLogger = logrus.New()
}

// SetLogger receives a logger and stores it as the base logger for the package.
func SetLogger(logger logrus.FieldLogger) {
	baseLogger = logger
}

// FromContext returns an logger with all values from context loaded on it.
func FromContext(ctx context.Context) logrus.FieldLogger {
	log := baseLogger
	if id, err := traceID(ctx); err == nil {
		log = log.WithField("traceID", id)
	}

	if id, err := messageID(ctx); err == nil {
		log = log.WithField("messageID", id)
	}

	return log
}

func traceID(ctx context.Context) (string, error) {
	traceID, ok := ctx.Value(TraceIDKey).(string)
	if !ok {
		return "", errors.New("No TraceID in the context")
	}

	return traceID, nil
}

func messageID(ctx context.Context) (string, error) {
	messageID, ok := ctx.Value(MessageIDKey).(string)
	if !ok {
		return "", errors.New("No MessageID in the context")
	}

	return messageID, nil
}
