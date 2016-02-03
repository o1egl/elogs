package logrus

import log "github.com/Sirupsen/logrus"

type logger struct {
	l *log.Logger
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

func New() *logger {
	return &logger{log.StandardLogger()}
}

func FromLogger(l *log.Logger) *logger {
	return &logger{l}
}
