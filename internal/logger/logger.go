package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

// Initialize sets up the logger with the specified log level and optional
// file logging. If logToFile is true, the logger will write to a file in
// addition to the standard output. The log file will be located in
// ~/.gopix/logs/gopix.log.
func Initialize(level string, logToFile bool) error {
	Logger = logrus.New()

	// Set log level
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	Logger.SetLevel(logLevel)

	// Set formatter
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	// Setup file logging if requested
	if logToFile {
		homeDir, _ := os.UserHomeDir()
		logDir := filepath.Join(homeDir, ".gopix", "logs")
		if err := os.MkdirAll(logDir, 0755); err == nil {
			logFile := filepath.Join(logDir, "gopix.log")
			file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				mw := io.MultiWriter(os.Stdout, file)
				Logger.SetOutput(mw)
			}
		}
	}

	return nil
}
