package logrus

import (
	log "github.com/Sirupsen/logrus"
	"io"
)

type logger struct {
	l *log.Logger
}

func (l *logger) SetOutput(w io.Writer) {
	l.l.Out = w
}

func (l *logger) SetLevel(lvl uint8) {
	switch lvl {
	case 0:
		l.l.Level = log.PanicLevel
	case 1:
		l.l.Level = log.FatalLevel
	case 2:
		l.l.Level = log.ErrorLevel
	case 3:
		l.l.Level = log.WarnLevel
	case 4:
		l.l.Level = log.InfoLevel
	default:
		l.l.Level = log.DebugLevel
	}
}

func (l *logger) Print(args ...interface{}) {
	l.l.Print(args...)
}
func (l *logger) Printf(format string, args ...interface{}) {
	l.l.Printf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.l.Debug(args...)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.l.Debugf(format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.l.Info(args...)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.l.Infof(format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.l.Warn(args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.l.Warnf(format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.l.Error(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.l.Errorf(format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.l.Fatal(args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.l.Fatalf(format, args...)
}

// New returns logrus implementation
func New() *logger {
	return &logger{log.StandardLogger()}
}

// FromLogger returns logrus implementation from custom logger instance
func FromLogger(l *log.Logger) *logger {
	return &logger{l}
}
