package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	FATAL

	Debug = "Debug"
	Info  = "Info"
	Warn  = "Warn"
	Error = "Error"
	Fatal = "Fatal"

	SHORT = iota
	FULL

	FilePathShort = "Short"
	FilePathFull  = "Full"
)

type Logger struct {
	l      *log.Logger
	config *Config
}

func New(config *Config) *Logger {
	l := &Logger{}
	l.config = config
	l.Init()
	go l.SendLogs()
	return l
}

func (l *Logger) Init() error {
	l.l = log.New(os.Stdout, fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference), 0)
	return nil
}

func (l *Logger) GetRef() string {
	return l.config.Reference
}

func (l *Logger) Debug(v ...interface{}) {
	if l.config.Level > DEBUG {
		return
	}

	if l.config.remoteLogger {
		go l.PostToRemote("DEBUG", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintln(l.formatLog("DEBUG", v...)...))
		return
	}

	l.l.Println(l.formatLog("DEBUG", v...)...)
}

func (l *Logger) Info(v ...interface{}) {
	if l.config.Level > INFO {
		return
	}

	if l.config.remoteLogger {
		go l.PostToRemote("INFO", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintln(l.formatLog("INFO", v...)...))
		return
	}

	l.l.Println(l.formatLog("INFO", v...)...)
}

func (l *Logger) Warn(v ...interface{}) {
	if l.config.Level > WARN {
		return
	}

	if l.config.remoteLogger {
		go l.PostToRemote("WARN", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintln(l.formatLog("WARN", v...)...))
		return
	}

	l.l.Println(l.formatLog("WARN", v...)...)
}

func (l *Logger) Error(v ...interface{}) {
	if l.config.Level > ERROR {
		return
	}

	if l.config.remoteLogger {
		go l.PostToRemote("ERROR", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintln(l.formatLog("ERROR", v...)...))
		return
	}

	l.l.Println(l.formatLog("ERROR", v...)...)
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.config.Level > FATAL {
		return
	}

	if l.config.remoteLogger {
		go l.PostToRemote("FATAL", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintln(l.formatLog("FATAL", v...)...))
		return
	}

	l.l.Println(l.formatLog("FATAL", v...)...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.config.Level > DEBUG {
		return
	}

	format, v = l.formatLogf("DEBUG", format, v...)
	if l.config.remoteLogger {
		l.PostToRemote("DEBUG", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintf(format, v...))
		return
	}

	l.l.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.config.Level > INFO {
		return
	}

	format, v = l.formatLogf("INFO", format, v...)
	if l.config.remoteLogger {
		go l.PostToRemote("INFO", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintf(format, v...))
		return
	}

	l.l.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.config.Level > WARN {
		return
	}

	format, v = l.formatLogf("WARN", format, v...)
	if l.config.remoteLogger {
		go l.PostToRemote("WARN", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintf(format, v...))
		return
	}

	l.l.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.config.Level > ERROR {
		return
	}

	format, v = l.formatLogf("ERROR", format, v...)
	if l.config.remoteLogger {
		go l.PostToRemote("ERROR", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintf(format, v...))
		return
	}

	l.l.Printf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.config.Level > FATAL {
		return
	}

	format, v = l.formatLogf("FATAL", format, v...)
	if l.config.remoteLogger {
		go l.PostToRemote("FATAL", fmt.Sprintf("%v [%s] [ %s ] ", time.Now().UTC(), l.config.AppName, l.config.Reference)+fmt.Sprintf(format, v...))
		return
	}

	l.l.Printf(format, v...)
}

// Format the log to contain the log levels
func (l *Logger) formatLog(logType string, v ...interface{}) []interface{} {
	var n []interface{}
	n = append(n, "["+logType+"] ")
	file, line := l.GetFileLine(3)

	n = append(n, file+":"+strconv.Itoa(line)+":")
	n = append(n, v...)
	return n
}

// Format the log to contain the log levels
func (l *Logger) formatLogf(logType string, format string, v ...interface{}) (string, []interface{}) {
	var n []interface{}
	prefix := "[%s] "
	n = append(n, logType)
	file, line := l.GetFileLine(3)

	prefix += "%s:%d: "
	format = prefix + format
	n = append(n, file)
	n = append(n, line)
	n = append(n, v...)
	return format, n
}

func (l *Logger) GetFileLine(n int) (string, int) {
	_, file, line, _ := runtime.Caller(n)
	// If you want the short path not the full file path, you can uncomment everything below
	if l.config.FilePathSize == SHORT {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
	}
	return file, line
}
