package gologger

import (
	"fmt"
	"github.com/apsdehal/go-logger"
)

type golog struct {
	l *logger.Logger
}

func (l *golog) Debug(args ...interface{}) {
	l.l.Debug(fmt.Sprint(args...))
}
func (l *golog) Debugf(format string, args ...interface{}) {
	l.l.Debug(fmt.Sprintf(format, args...))
}

func (l *golog) Info(args ...interface{}) {
	l.l.Info(fmt.Sprint(args...))
}
func (l *golog) Infof(format string, args ...interface{}) {
	l.l.Info(fmt.Sprintf(format, args...))
}

func (l *golog) Warn(args ...interface{}) {
	l.l.Warning(fmt.Sprint(args...))
}

func (l *golog) Warnf(format string, args ...interface{}) {
	l.l.Warning(fmt.Sprintf(format, args...))
}

func (l *golog) Error(args ...interface{}) {
	l.l.Error(fmt.Sprint(args...))
}

func (l *golog) Errorf(format string, args ...interface{}) {
	l.l.Error(fmt.Sprintf(format, args...))
}

func (l *golog) Fatal(args ...interface{}) {
	l.l.Fatal(fmt.Sprint(args...))
}

func (l *golog) Fatalf(format string, args ...interface{}) {
	l.l.Fatal(fmt.Sprintf(format, args...))
}

// FromLogger returns go-logger implementation
func New() *golog {
	l, _ := logger.New()
	return &golog{l}
}

// FromLogger returns go-logger implementation from custom instance
func FlomLogger(l *logger.Logger) *golog {
	return &golog{l}
}
