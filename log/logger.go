package log

import (
	"bytes"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(b, "%s \033[%dm %s \033[0m %s %s: %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s \x1b[%dm %s \x1b[0m: %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()

	// Create log file writer
	fileWriter := &lumberjack.Logger{
		Filename:   "log/app.log",
		MaxSize:    10, // MB
		MaxBackups: 7,
		MaxAge:     30, // days
		Compress:   true,
	}

	// Create multi-writer to write to both stdout and log file
	multiWriter := io.MultiWriter(os.Stdout, fileWriter)

	mLog.SetOutput(multiWriter)        // Set output to both stdout and file
	mLog.SetReportCaller(true)         // Enable reporting of caller
	mLog.SetFormatter(&LogFormatter{}) // Set custom formatter

	level, err := logrus.ParseLevel("info")
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) // Set minimum log level
	InitDefaultLogger()  // Initialize default logger for global use
	return mLog
}

func InitDefaultLogger() {
	// Set global log output to both stdout and file
	fileWriter := &lumberjack.Logger{
		Filename:   "log/app.log",
		MaxSize:    10, // MB
		MaxBackups: 7,
		MaxAge:     30, // days
		Compress:   true,
	}

	multiWriter := io.MultiWriter(os.Stdout, fileWriter)

	logrus.SetOutput(multiWriter)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel("info")
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
