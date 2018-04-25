package log

import (
	"github.com/Sirupsen/logrus"
)

// Logger variable
var Logger *logrus.Logger

// Init func
func Init() {
	Logger = logrus.New()
}

func New() *logrus.Logger {
	if Logger != nil {
		return Logger
	}
	return logrus.New()
}
