package glog

import (
	"github.com/golang/glog"
	"io"
)

type logger struct {
}

func (l *logger) SetOutput(w io.Writer) {
}

func (l *logger) SetLevel(lvl uint8) {
}

func (l *logger) Print(args ...interface{}) {
	glog.Info(args...)
}
func (l *logger) Printf(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (*logger) Debug(args ...interface{}) {
	glog.Info(args...)
}
func (*logger) Debugf(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (*logger) Info(args ...interface{}) {
	glog.Info(args...)
}
func (*logger) Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (*logger) Warn(args ...interface{}) {
	glog.Warning(args...)
}

func (*logger) Warnf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func (*logger) Error(args ...interface{}) {
	glog.Error(args...)
}

func (*logger) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func (*logger) Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func (*logger) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}

// New returns glog implementation
func New() *logger {
	return &logger{}
}
