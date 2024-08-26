package util

import (
	"flag"
	"management/model"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	Logger.Out = os.Stdout
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,  // Forces the use of colors even if the output is not a terminal
		FullTimestamp: false, // Adds timestamps to logs
	})
}

func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level (debug, error, info, warning)")

	flag.Parse()

	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func Log(logLevel, packageLevel, functionName string, message, parameter interface{}) {
	switch logLevel {
	case model.LogLevelDebug:
		if parameter != nil {
			Logger.Debug("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Debug("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)

		}
	case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)

		}
	case model.LogLevelWarning:
		if parameter != nil {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)

		}
	default:
		if parameter != nil {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)

		}
	}
}
