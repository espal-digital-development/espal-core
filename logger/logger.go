package logger

import (
	"fmt"
	"sync"
)

// TODO :: 77 Add config option for writing `timestamps` in the message and implement it in here

var _ Loggable = &Logger{}

// Types identify what style the logger is meant for.
const (
	TypeInfo uint8 = iota + 1
	TypeWarning
	TypeError
	TypeCustom
)

// logMessage holds the message information that get send to the logger instance.
type logMessage struct {
	_type          uint8
	message        string
	customCallback func(string) string
}

// Loggable represents and object that can consume messages of different types
// representing different levels of urgency, which then get processed to be
// written or streamed to an internally defined destination.
type Loggable interface {
	Disable()
	Enable()
	IsEnabled() bool
	Info(message string)
	Infof(message string, params ...interface{})
	Warning(message string)
	Warningf(message string, params ...interface{})
	Error(message string)
	Errorf(message string, params ...interface{})
	Custom(message string, callback func(string) string)
	Customf(message string, callback func(string) string, params ...interface{})
	GetLastMessage() string
	EnableTestMode()
}

// Logger are used to handle non-blocking streams of log outputs.
type Logger struct {
	// TODO :: The messages should be able to be logged locally, but also send to an outside resource/endpoint
	messages chan logMessage
	enabled  bool

	lastMessage string
	testMutex   *sync.RWMutex
	testMode    bool
}

// Disable disables logging going outward.
func (l *Logger) Disable() {
	l.enabled = false
}

// Enable enables logging going outward.
func (l *Logger) Enable() {
	l.enabled = true
}

// IsEnabled indicates if the loggin is currently active.
func (l *Logger) IsEnabled() bool {
	return l.enabled
}

// Info logs a message to the Logger of the type TypeInfo.
func (l *Logger) Info(message string) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:   TypeInfo,
		message: message,
	}
}

// Infof logs a formatted message to the Logger of the type TypeInfo.
func (l *Logger) Infof(message string, params ...interface{}) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:   TypeInfo,
		message: fmt.Sprintf(message, params...),
	}
}

// Warning logs a message to the Logger of the type TypeWarning.
func (l *Logger) Warning(message string) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:   TypeWarning,
		message: message,
	}
}

// Warningf logs a formatted message to the Logger of the type TypeWarning.
func (l *Logger) Warningf(message string, params ...interface{}) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:   TypeWarning,
		message: fmt.Sprintf(message, params...),
	}
}

// Error logs a message to the Logger of the type TypeError.
func (l *Logger) Error(message string) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:   TypeError,
		message: message,
	}
}

// Errorf logs a formatted message to the Logger of the type TypeError.
func (l *Logger) Errorf(message string, params ...interface{}) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:   TypeError,
		message: fmt.Sprintf(message, params...),
	}
}

// Custom logs a message to the Logger of the type TypeCustom.
func (l *Logger) Custom(message string, callback func(string) string) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:          TypeCustom,
		message:        message,
		customCallback: callback,
	}
}

// Customf logs a message to the Logger of the type TypeCustom.
func (l *Logger) Customf(message string, callback func(string) string, params ...interface{}) {
	if !l.enabled {
		return
	}
	l.messages <- logMessage{
		_type:          TypeCustom,
		message:        fmt.Sprintf(message, params...),
		customCallback: callback,
	}
}

// GetLastMessage returns the last message that was logged.
// This only works when testMode is active. Don't use in production.
func (l *Logger) GetLastMessage() string {
	l.testMutex.RLock()
	defer l.testMutex.RUnlock()
	return l.lastMessage
}

func (l *Logger) setLastMessage(message string) {
	l.testMutex.Lock()
	defer l.testMutex.Unlock()
	l.lastMessage = message
}

// EnableTestMode enables test mode for tests interaction.
func (l *Logger) EnableTestMode() {
	l.testMode = true
}

func (l *Logger) startListener(listener func(*Logger)) {
	go listener(l)
}

// New returns a new pointer-instance of Logger.
func New() *Logger {
	l := &Logger{
		messages:  make(chan logMessage),
		testMutex: &sync.RWMutex{},
	}
	l.Enable()
	l.startListener(func(logger *Logger) {
		for logger != nil {
			message := <-logger.messages
			switch message._type {
			case TypeInfo:
				fmt.Println("INFO  :", message.message)
				if l.testMode {
					logger.setLastMessage(fmt.Sprint("INFO  : ", message.message))
				}
			case TypeWarning:
				fmt.Println("WARN  :", message.message)
				if l.testMode {
					logger.setLastMessage(fmt.Sprint("WARN  : ", message.message))
				}
			case TypeError:
				fmt.Println("ERROR :", message.message)
				if l.testMode {
					logger.setLastMessage(fmt.Sprint("ERROR : ", message.message))
				}
			case TypeCustom:
				fmt.Println(message.customCallback(message.message))
				if l.testMode {
					logger.setLastMessage(fmt.Sprint(message.customCallback(message.message)))
				}
			default:
				fmt.Printf("Unknown Logger Message Type `%d`\n", message._type)
				if l.testMode {
					logger.setLastMessage(fmt.Sprintf("Unknown Logger Message Type `%d`\n", message._type))
				}
			}
		}
	})
	return l
}
