package logger

import (
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

 func LoadLogger(level string) error{

	switch level{
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.DateTime,
		})
		return nil

	case "PRODUCTION":
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.DateTime,
		})
		file, err := os.OpenFile("../logs/subsciption.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    	if err == nil {
			mw := io.MultiWriter(os.Stdout, file)
        	logrus.SetOutput(mw)
		}
		return err
	
	default: 
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.DateTime,
		})
		return nil
	}
}