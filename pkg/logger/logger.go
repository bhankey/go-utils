package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

// GetLogger Init initialize logger.
func GetLogger(level int) (Logger, error) {
	log := logrus.New()

	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: false,
	})

	log.SetLevel(logrus.Level(level))

	return Logger{log}, nil
}
