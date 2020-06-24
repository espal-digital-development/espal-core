package logger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/espal-digital-development/espal-core/logger"
)

const (
	param1       = "custom 1"
	param2       = "custom 2"
	errorMessage = "I'm an error message"
)

func expectDelayedMessage(t *testing.T, logger logger.Loggable, expectedMessage string) {
	time.Sleep(time.Millisecond)
	loggerMessage := logger.GetLastMessage()
	if expectedMessage != loggerMessage {
		t.Errorf("expected logger to give `%s`, but got `%s`", expectedMessage, loggerMessage)
	}
}

func TestNew(t *testing.T) {
	logger := logger.New()
	if logger == nil {
		t.Fatal("expected logger to not be nil")
	}
}

func TestSwitchState(t *testing.T) {
	logger := logger.New()
	if !logger.IsEnabled() {
		t.Error("the initial state of a Logger should be enabled")
	}
	logger.Enable()
	if !logger.IsEnabled() {
		t.Fatal("failed to enable the Logger")
	}
	logger.Disable()
	if logger.IsEnabled() {
		t.Fatal("failed to disable the Logger")
	}
}

func TestInfoEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	message := "I'm an info message"
	logger.Info(message)
	expectDelayedMessage(t, logger, "INFO  : "+message)
}

func TestInfoDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an info message"
	logger.Info(message)
	expectDelayedMessage(t, logger, "")
}

func TestInfofEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	message := "I'm an %s %s info message"
	logger.Infof(message, param1, param2)
	expectDelayedMessage(t, logger, "INFO  : "+fmt.Sprintf(message, param1, param2))
}

func TestInfofDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an %s %s info message"
	logger.Infof(message, param1, param2)
	expectDelayedMessage(t, logger, "")
}

func TestWarningEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	message := "I'm an warning message"
	logger.Warning(message)
	expectDelayedMessage(t, logger, "WARN  : "+message)
}

func TestWarningDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an warning message"
	logger.Warning(message)
	expectDelayedMessage(t, logger, "")
}

func TestWarningfEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	message := "I'm an %s %s warning message"
	logger.Warningf(message, param1, param2)
	expectDelayedMessage(t, logger, "WARN  : "+fmt.Sprintf(message, param1, param2))
}

func TestWarningfDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an %s %s warning message"
	logger.Warningf(message, param1, param2)
	expectDelayedMessage(t, logger, "")
}

func TestErrorEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Error(errorMessage)
	expectDelayedMessage(t, logger, "ERROR : "+errorMessage)
}

func TestErrorDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	logger.Error(errorMessage)
	expectDelayedMessage(t, logger, "")
}

func TestErrorfEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	message := "I'm an %s %s error message"
	logger.Errorf(message, param1, param2)
	expectDelayedMessage(t, logger, "ERROR : "+fmt.Sprintf(message, param1, param2))
}

func TestErrorfDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an %s %s error message"
	logger.Errorf(message, param1, param2)
	expectDelayedMessage(t, logger, "")
}

func TestCustomEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Custom(errorMessage, func(s string) string {
		return s
	})
	expectDelayedMessage(t, logger, errorMessage)
}

func TestCustomDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an custom message"
	logger.Custom(message, func(s string) string {
		return s
	})
	expectDelayedMessage(t, logger, "")
}

func TestCustomfEnabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	message := "I'm an %s %s message"
	param1 := "custom 1"
	param2 := "custom 2"
	logger.Customf(message, func(s string) string {
		return s
	}, param1, param2)
	expectDelayedMessage(t, logger, fmt.Sprintf(message, param1, param2))
}
func TestCustomfDisabled(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an %s %s message"
	param1 := "custom 1"
	param2 := "custom 2"
	logger.Customf(message, func(s string) string {
		return s
	}, param1, param2)
	expectDelayedMessage(t, logger, "")
}

func TestDisabledAndReEnable(t *testing.T) {
	logger := logger.New()
	logger.EnableTestMode()
	logger.Disable()
	message := "I'm an error message"
	logger.Error(message)
	expectDelayedMessage(t, logger, "")
	logger.Enable()
	logger.Error(message)
	expectDelayedMessage(t, logger, "ERROR : "+message)
}
