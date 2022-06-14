package logger

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

type LogMessage struct {
	component string
	level     logrus.Level
	message   string
}

var COMPONENT = "component"

type LoggerWrapper struct {
	stdLogger   *logrus.Logger
	fileLogger  *logrus.Logger
	file        *os.File
	serviceName string
	scheduler   *gocron.Scheduler
}

func InitLogger(serviceName string, s *gocron.Scheduler) *LoggerWrapper {
	file := openFile(generateLogFileName(serviceName))
	stdLogger := logrus.New()
	stdLogger.Out = os.Stdout
	fileLogger := logrus.New()
	fileLogger.Out = file
	fileLogger.SetFormatter(&logrus.JSONFormatter{})
	stdLogger.SetFormatter(&logrus.JSONFormatter{})
	loggerWrap := &LoggerWrapper{stdLogger: stdLogger, fileLogger: fileLogger, file: file, serviceName: serviceName, scheduler: s}
	err := s.Every(24).Hour().Do(loggerWrap.rotate)
	if err != nil {
		return nil
	}
	return loggerWrap
}

func (logger *LoggerWrapper) Writeln(message LogMessage) {
	logger.fileLogger.WithFields(logrus.Fields{
		COMPONENT: message.component,
	}).Log(message.level, message.message)

	logger.stdLogger.WithFields(logrus.Fields{
		COMPONENT: message.component,
	}).Log(message.level, message.message)
}

func (logger *LoggerWrapper) rotate() {
	file := openFile(generateLogFileName(logger.serviceName))
	logger.file.Close()
	logger.file = file
	logger.fileLogger.Out = file
}
func generateLogFileName(serviceName string) string {
	filename := fmt.Sprintf("logs_%s_%s.log", strconv.FormatInt(time.Now().Unix(), 10), serviceName)
	return filename
}
func openFile(filename string) *os.File {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	return file
}
