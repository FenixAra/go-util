package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

type Logger struct {
	l   *log.Logger
	ref string
}

func New(ref string) *Logger {
	l := &Logger{}
	l.Init(ref)
	return l
}

func (l *Logger) Init(ref string) error {
	l.ref = ref
	if l.ref == "" {
		refuuid, err := uuid.NewV4()
		if err != nil {
			l.Error("Unable to generate new UUID. Err: ", err)
			return err
		}

		l.ref = refuuid.String()
	}
	l.l = log.New(os.Stdout, fmt.Sprintf("[ %s ] ", l.ref), 0)
	return nil
}

func (l *Logger) GetRef() string {
	return l.ref
}

func (l *Logger) Warn(a ...interface{}) {
	l.l.Println(formatLog("WARN", a...)...)
}

func (l *Logger) Error(a ...interface{}) {
	l.l.Println(formatLog("ERROR", a...)...)
}

func (l *Logger) Info(a ...interface{}) {
	l.l.Println(formatLog("INFO", a...)...)
}

func (l *Logger) Debug(a ...interface{}) {
	l.l.Println(formatLog("DEBUG", a...)...)
}

func (l *Logger) Fatal(a ...interface{}) {
	l.l.Println(formatLog("FATAL", a...)...)
}

// Format the log to contain the log levels
func formatLog(logType string, a ...interface{}) []interface{} {
	var n []interface{}
	n = append(n, "["+logType+"] ")
	_, file, line, _ := runtime.Caller(2)
	// If you want the short path not the full file path, you can uncomment everything below
	/*
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
	*/

	n = append(n, file+":"+strconv.Itoa(line)+":")
	n = append(n, a...)
	return n
}
