package log

import (
	"github.com/miiren/mbox/logrus"
)

var (
	Logger       logrus.ULogger
	LoggerAccess logrus.ULogger
	LoggerDb     logrus.ULogger
)

func InitLogger(filepath string) {
	Logger = logrus.ULogger{Logger: logrus.NewDefaultLogger(filepath)}
	LoggerAccess = logrus.ULogger{Logger: logrus.New(filepath, "access")}
	LoggerDb = logrus.ULogger{Logger: logrus.New(filepath, "orm")}
}
