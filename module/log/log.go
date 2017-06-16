package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	level          logrus.Level
	file           *logrus.Logger
	logConsoleFlag bool //log to console if flag is true.
}

var logger Logger

func init() {
	logger.level = logrus.ErrorLevel
	logger.file = logrus.StandardLogger()
	logger.file.Level = logger.level
	logger.logConsoleFlag = false
}

func InitLogger(logLvStr string, dir string, file string, logConsoleFlag bool,
) error {
	level, err := logrus.ParseLevel(logLvStr)
	if err != nil {
		return err
	}
	logger.level = level

	dir = strings.TrimRight(dir, " /")
	err = os.MkdirAll(dir, os.FileMode(0777))
	if err != nil {
		return err
	}

	//log file name format: dir/file_year_month_day_UnixTimestamp.log
	t := time.Now()
	tSuffix := fmt.Sprintf("_%d_%02d_%02d_%d", t.Year(), t.Month(), t.Day(),
		t.Unix())
	filename := dir + "/" + file + tSuffix + ".log"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	logger.file = logrus.New()
	logger.file.Out = f
	logger.file.Level = level
	logger.file.Formatter = &logrus.JSONFormatter{}
	runtime.SetFinalizer(logger.file, func(*logrus.Logger) {
		file, ok := logger.file.Out.(*os.File)
		if ok {
			err := file.Close()
			if err != nil {
				logrus.Errorf("Close log file error: %v.", err)
			}
		}
	})

	logger.logConsoleFlag = logConsoleFlag
	logrus.SetLevel(level)

	return nil
}

func GetLogger() *Logger {
	return &logger
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.file.Debugf(format, args...)
	if l.logConsoleFlag {
		logrus.Debugf(format, args...)
	}
}
func (l *Logger) Info(format string, args ...interface{}) {
	l.file.Infof(format, args...)
	if l.logConsoleFlag {
		logrus.Infof(format, args...)
	}
}
func (l *Logger) Warn(format string, args ...interface{}) {

	l.file.Warnf(format, args...)
	if l.logConsoleFlag {
		logrus.Warnf(format, args...)
	}
}
func (l *Logger) Error(format string, args ...interface{}) {

	l.file.Errorf(format, args...)
	if l.logConsoleFlag {
		logrus.Errorf(format, args...)
	}
}
func (l *Logger) Fatal(format string, args ...interface{}) {

	l.file.Fatalf(format, args...)
	if l.logConsoleFlag {
		logrus.Fatalf(format, args...)
	}
}
func (l *Logger) Panic(format string, args ...interface{}) {

	l.file.Panicf(format, args...)
	if l.logConsoleFlag {
		logrus.Panicf(format, args...)
	}
}
func (l *Logger) SetLogLevel(lvStr string) error {
	level, err := logrus.ParseLevel(lvStr)
	if err != nil {
		return err
	}
	l.level = level
	l.file.Level = level
	logrus.SetLevel(level)
	return nil
}
func (l *Logger) SetLogConsole() {
	l.logConsoleFlag = true
}
func (l *Logger) CloseLogConsole() {
	l.logConsoleFlag = false
}
