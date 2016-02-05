package glog

import (
	"github.com/labstack/echo"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	Convey("Logger should", t, func() {
		Convey("Implement echo.Logger interface", func() {
			l := New()
			So(l, ShouldImplement, (*echo.Logger)(nil))
		})
	})
}

func TestGlogLogger(t *testing.T) {
	l := New()
	switch os.Getenv("TEST_FATAL") {
	case "1":
		l.Fatal("Fatal")
	case "2":
		l.Fatalf("Fatal%s", "f")
	}
	Convey("Logger should print", t, func() {
		Convey("Debug", func() {
			l.Debug("Debug")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Debugf", func() {
			l.Debugf("Debug%s", "f")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Info", func() {
			l.Info("Info")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Infof", func() {
			l.Infof("Info%s", "f")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Warn", func() {
			l.Warn("Warn")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Warnf", func() {
			l.Warnf("Warn%s", "f")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Error", func() {
			l.Error("Error")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
		Convey("Errorf", func() {
			l.Errorf("Error%s", "f")
			// TODO: Check message
			So(true, ShouldBeTrue)
		})
	})
}
