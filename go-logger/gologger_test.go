package gologger

import (
	"bytes"
	"github.com/apsdehal/go-logger"
	"github.com/labstack/echo/log"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	Convey("Logger should", t, func() {
		Convey("Implement echo.Logger interface", func() {
			l := New()
			So(l, ShouldImplement, (*log.Logger)(nil))
		})
		Convey("Initialized from custom logger", func() {
			buf := new(bytes.Buffer)
			logger, _ := logger.New("test", 0, buf)
			l := FlomLogger(logger)
			l.Debug("Debug")
			So(strings.Contains(buf.String(), "Debug"), ShouldBeTrue)
		})
	})
}

func TestLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	logger, _ := logger.New("test", 0, buf)
	l := FlomLogger(logger)
	l.SetLevel(10)
	l.SetOutput(buf)
	Convey("Logger should print", t, func() {
		Convey("Print", func() {
			l.Print("Print")
			testLogMsg("Print", buf)
		})
		Convey("Printf", func() {
			l.Printf("Print%s", "f")
			testLogMsg("Printf", buf)
		})
		Convey("Debug", func() {
			l.Debug("Debug")
			testLogMsg("Debug", buf)
		})
		Convey("Debugf", func() {
			l.Debugf("Debug%s", "f")
			testLogMsg("Debugf", buf)
		})
		Convey("Info", func() {
			l.Info("Info")
			testLogMsg("Info", buf)
		})
		Convey("Infof", func() {
			l.Infof("Info%s", "f")
			testLogMsg("Infof", buf)
		})
		Convey("Warn", func() {
			l.Warn("Warn")
			testLogMsg("Warn", buf)
		})
		Convey("Warnf", func() {
			l.Warnf("Warn%s", "f")
			testLogMsg("Warnf", buf)
		})
		Convey("Error", func() {
			l.Error("Error")
			testLogMsg("Error", buf)
		})
		Convey("Errorf", func() {
			l.Errorf("Error%s", "f")
			testLogMsg("Errorf", buf)
		})
		Convey("Fatal", func() {
			cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
			cmd.Env = append(os.Environ(), "TEST_FATAL=1")
			cmd.Stdout = buf
			cmd.Stderr = buf
			cmd.Run()
			testLogMsg("Fatal", buf)
		})
		Convey("Fatalf", func() {
			cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
			cmd.Env = append(os.Environ(), "TEST_FATAL=2")
			cmd.Stdout = buf
			cmd.Stderr = buf
			cmd.Run()
			testLogMsg("Fatalf", buf)
		})
	})
}

func TestFatal(t *testing.T) {
	l := New()
	switch os.Getenv("TEST_FATAL") {
	case "1":
		l.Fatal("Fatal")
	case "2":
		l.Fatalf("Fatal%s", "f")
	}
}

func testLogMsg(msg string, buf *bytes.Buffer) {
	So(strings.Contains(buf.String(), msg), ShouldBeTrue)
	buf.Reset()
}
