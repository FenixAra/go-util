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
	Level     string    `json:"level"`
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"message"`
	Message   string    `json:"full_message"`
	AppName   string    `json:"app_name"`
	RefID     string    `json:"ref_id"`
	File      string    `json:"file"`
	Line      string    `json:"line"`
}

func (l *Logger) PostToRemote(level, msg string) {
	client := &http.Client{}
	file, line := l.GetFileLine(3)
	reqData, err := json.Marshal(&Log{
		Level:     level,
		Timestamp: time.Now().UTC(),
		Title:     msg,
		Message:   msg,
		AppName:   l.config.AppName,
		RefID:     l.config.Reference,
		File:      file,
		Line:      strconv.Itoa(line),
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
