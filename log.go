package log

import (
	"fmt"
	"io"
	golog "log"
	"strings"

	"os"
)

// Logger level
const (
	LvOff = iota
	LvCritical
	LvError
	LvWarn
	LvInfo
	LvDebug
	LvTrace
	LvAll
)

// ToLevel ...
func ToLevel(level string) int {
	switch strings.ToLower(level) {
	case "off":
		return LvOff
	case "critical":
		return LvCritical
	case "error":
		return LvError
	case "warn":
		return LvWarn
	case "info":
		return LvInfo
	case "debug":
		return LvDebug
	case "trace":
		return LvTrace
	default:
		return LvAll
	}
}

// Logger ...
type Logger interface {
	Trace(v ...interface{})
	Tracef(string, ...interface{})

	Debug(v ...interface{})
	Debugf(string, ...interface{})

	Info(v ...interface{})
	Infof(string, ...interface{})

	Warn(v ...interface{})
	Warnf(string, ...interface{})

	Error(v ...interface{})
	Errorf(string, ...interface{})

	Critical(v ...interface{})
	Criticalf(string, ...interface{})
}

// global variables
const (
	Flags = golog.Lshortfile | golog.LstdFlags | golog.Lmicroseconds /*| golog.LUTC*/
)

var (
	loggers = make(map[string]Logger)
	root    = New(os.Stderr, LvAll, 4)
	root3   = New(os.Stderr, LvAll, 3)
)

type _logger struct {
	*golog.Logger
	level int
	deep  int
}

// Register register logger
func Register(name string, l Logger) {
	if name == "root" {
		root = l
	} else {
		loggers[name] = l
	}
}

// New return new Logger
func New(w io.Writer, level, deep int) Logger {
	l := golog.New(w, "", Flags)
	ret := &_logger{
		Logger: l,
		level:  level,
		deep:   deep,
	}

	return ret
}

// Get return a logger
func Get(name string) Logger {
	ret, ok := loggers[name]
	if !ok || ret == nil {
		Warnf("can find %s, return default", name)
		return root3
	}
	return ret
}

func (x *_logger) logf(prefix string, format string, v ...interface{}) {
	x.Output(x.deep, "["+prefix+"] "+fmt.Sprintf(format, v...))
}

func (x *_logger) log(prefix string, v ...interface{}) {
	x.Output(x.deep, "["+prefix+"] "+fmt.Sprint(v...))
}

func (x *_logger) Trace(v ...interface{}) {
	if x.level >= LvTrace {
		x.log("TRACE", v...)
	}
}

func (x *_logger) Tracef(format string, v ...interface{}) {
	if x.level >= LvTrace {
		x.logf("TRACE", format, v...)
	}
}

func (x *_logger) Debug(v ...interface{}) {
	if x.level >= LvDebug {
		x.log("DEBUG", v...)
	}
}

func (x *_logger) Debugf(format string, v ...interface{}) {
	if x.level >= LvDebug {
		x.logf("DEBUG", format, v...)
	}
}

func (x *_logger) Info(v ...interface{}) {
	if x.level >= LvInfo {
		x.log("INFO", v...)
	}
}

func (x *_logger) Infof(format string, v ...interface{}) {
	if x.level >= LvInfo {
		x.logf("INFO", format, v...)
	}
}

func (x *_logger) Warn(v ...interface{}) {
	if x.level >= LvWarn {
		x.log("WARN", v...)
	}
}

func (x *_logger) Warnf(format string, v ...interface{}) {
	if x.level >= LvWarn {
		x.logf("WARN", format, v...)
	}
}

func (x *_logger) Error(v ...interface{}) {
	if x.level >= LvError {
		x.log("ERROR", v...)
	}
}

func (x *_logger) Errorf(format string, v ...interface{}) {
	if x.level >= LvError {
		x.logf("ERROR", format, v...)
	}
}

func (x *_logger) Critical(v ...interface{}) {
	if x.level >= LvCritical {
		x.log("CRITICAL", v...)
	}
}

func (x *_logger) Criticalf(format string, v ...interface{}) {
	if x.level >= LvCritical {
		x.logf("CRITICAL", format, v...)
	}
}

// Trace ...
func Trace(v ...interface{}) {
	root.Trace(v...)
}

// Tracef ...
func Tracef(format string, v ...interface{}) {
	root.Tracef(format, v...)
}

// Debug ...
func Debug(v ...interface{}) {
	root.Debug(v...)
}

// Debugf ...
func Debugf(format string, v ...interface{}) {
	root.Debugf(format, v...)
}

// Info ...
func Info(v ...interface{}) {
	root.Info(v...)
}

// Infof ...
func Infof(format string, v ...interface{}) {
	root.Infof(format, v...)
}

// Warn ...
func Warn(v ...interface{}) {
	root.Warn(v...)
}

// Warnf ...
func Warnf(format string, v ...interface{}) {
	root.Warnf(format, v...)
}

// Error ...
func Error(v ...interface{}) {
	root.Error(v...)
}

// Errorf ...
func Errorf(format string, v ...interface{}) {
	root.Errorf(format, v...)
}

// Critical ...
func Critical(v ...interface{}) {
	root.Critical(v...)
}

// Criticalf ...
func Criticalf(format string, v ...interface{}) {
	root.Criticalf(format, v...)
}
