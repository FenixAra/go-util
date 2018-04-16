package log

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Log struct {
	Level        string    `json:"level"`
	Timestamp    time.Time `json:"timestamp"`
	Title        string    `json:"message"`
	Message      string    `json:"full_message"`
	AppName      string    `json:"app_name"`
	RefID        string    `json:"ref_id"`
	File         string    `json:"file"`
	Line         string    `json:"line"`
	ResponseTime float64   `json:"response_time"`
	StatusCode   int       `json:"status_code"`
	Method       string    `json:"method"`
	Request      string    `json:"request"`
}

func (l *Logger) LogAPIInfo(request, method string, responseTime float64, status int) {
	if !l.config.RemoteLogger {
		return
	}

	l.postToRemote("INFO", "API Request info", request, method, responseTime, status)
}

func (l *Logger) postToRemote(level, msg, req, method string, responseTime float64, status int) {
	client := &http.Client{}
	file, line := l.GetFileLine(3)
	reqData, err := json.Marshal(&Log{
		Level:        level,
		Timestamp:    time.Now().UTC(),
		Title:        msg,
		Message:      msg,
		AppName:      l.config.AppName,
		RefID:        l.config.Reference,
		File:         file,
		Line:         strconv.Itoa(line),
		ResponseTime: responseTime,
		StatusCode:   status,
		Method:       method,
		Request:      req,
	})
	if err != nil {
		log.Println("Unable to marshal log request. Err:", err)
		return
	}

	request, err := http.NewRequest(http.MethodPost, l.config.RemoteLoggerURL, bytes.NewBuffer(reqData))
	if err != nil {
		log.Println("Unable to create new log request. Err:", err)
		return
	}

	// Setting all headers
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(request)
	if err != nil {
		log.Println("Unable to send log request. Err:", err)
		return
	}

	if response.StatusCode != http.StatusOK {
		log.Println("Status code is not 200 (OK). Got:", response.StatusCode)
		// Handle the error codes
		return
	}

	return
}

func (l *Logger) PostToRemote(level, msg string) {
	l.postToRemote(level, msg, "", "", 0, 0)
}
