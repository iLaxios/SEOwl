package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger(debug bool) {
	log = logrus.New()

	if debug {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func GetLogger() *logrus.Logger {
	return log
}
