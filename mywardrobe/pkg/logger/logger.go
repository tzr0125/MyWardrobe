package logger

import "github.com/sirupsen/logrus"

var Logger = logrus.StandardLogger()

func StdLogger() *logrus.Logger {
	return Logger
}
