package utils

import (
    "os"

    "github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func SetupLogger(level string) {
    logger = logrus.New()
    logger.Out = os.Stdout

    // Set the log level
    switch level {
    case "debug":
        logger.SetLevel(logrus.DebugLevel)
    case "info":
        logger.SetLevel(logrus.InfoLevel)
    case "warn":
        logger.SetLevel(logrus.WarnLevel)
    case "error":
        logger.SetLevel(logrus.ErrorLevel)
    default:
        logger.SetLevel(logrus.InfoLevel)
    }

    // Set the formatter
    logger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
}

func GetLogger() *logrus.Logger {
    if logger == nil {
        SetupLogger("info")
    }
    return logger
}
