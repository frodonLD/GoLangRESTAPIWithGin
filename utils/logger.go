package utils

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggingContext host our custom logging definition
type LoggingContext struct {
	LoggerEntry *logrus.Entry
	Level       logrus.Level
	Formatter   logrus.Formatter
	Output      *os.File
}

const (
	CorrelationIDKey  = "correlation_id"
	LoggingContextKey = "logging_context"
)

var (
	// Logger is a logrus entry that will be used as defauly logger
	Logger                *logrus.Entry
	defaultLoggingContext *LoggingContext
)

// NewLoggingContext -
func NewLoggingContext(level logrus.Level, formatter logrus.Formatter, output *os.File) *LoggingContext {
	lc := &LoggingContext{Level: level, Formatter: formatter, Output: output}
	logrusLogger := newLogger(lc.Level, lc.Output, lc.Formatter)
	lc.LoggerEntry = logrus.NewEntry(logrusLogger)
	return lc
}

// GetLoggingContextFromGinContext get a logginContext based on the gin context
func GetLoggingContextFromGinContext(c *gin.Context) *LoggingContext {
	if lc, ok := c.Get(LoggingContextKey); ok {
		return lc.(*LoggingContext)
	}
	return defaultLoggingContext
}

// SetCorrelationID add the correlatioID given as correlation id field to the logging context given
func (lc *LoggingContext) SetCorrelationID(correlationID string) {
	lc.LoggerEntry = lc.LoggerEntry.WithFields(logrus.Fields{
		CorrelationIDKey: correlationID,
	})
}

func newLogger(logLevel logrus.Level, output *os.File, formatter logrus.Formatter) *logrus.Logger {
	logger := logrus.New()
	logger.Level = logLevel
	logger.Formatter = formatter
	logger.Out = output
	return logger
}

func init() {
	defaultLoggingContext := NewLoggingContext(logrus.WarnLevel, &logrus.JSONFormatter{}, os.Stdout)
	Logger = defaultLoggingContext.LoggerEntry
}
