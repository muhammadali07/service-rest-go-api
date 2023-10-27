package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields

type Logger struct {
	log *logrus.Entry
}

func (l *Logger) Info(fields map[string]interface{}, method, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"method": method,
		}).WithFields(fields).Info(message)
	} else {
		l.log.WithFields(fields).Info(message)
	}
}
func (l *Logger) Warn(fields map[string]interface{}, method, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"method": method,
		}).WithFields(fields).Warn(message)
	} else {
		l.log.WithFields(fields).Warn(message)
	}
}
func (l *Logger) Error(fields map[string]interface{}, method string, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"method": method,
		}).WithFields(fields).Error(message)
	} else {
		l.log.WithFields(fields).Error(message)
	}
}
func (l *Logger) Fatal(fields map[string]interface{}, method string, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"method": method,
		}).WithFields(fields).Fatal(message)
	} else {
		l.log.WithFields(fields).Fatal(message)
	}
}
func (l *Logger) Panic(fields map[string]interface{}, method string, message string) {
	if method != "" {
		l.log.WithFields(logrus.Fields{
			"method": method,
		}).WithFields(fields).Panic(message)
	} else {
		l.log.WithFields(fields).Panic(message)
	}
}

func (l *Logger) SetFileOutput(file io.Writer) {
	l.log.Logger.Out = file
}

func NewLogger(name string) (log *Logger) {
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.FullTimestamp = true
	l := logrus.New()
	l.SetFormatter(formatter)
	log = &Logger{
		log: l.WithField("service", name),
	}
	return
}
