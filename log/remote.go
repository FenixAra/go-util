package log

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tomasen/realip"
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
	UserAgent    string    `json:"user_agent"`
	CustomerID   string    `json:"customer_id"`
	IPAddress    string    `json:"ip_address"`
}

func (l *Logger) LogAPIInfo(r *http.Request, responseTime float64, status int) {
	if !l.config.RemoteLogger {
		return
	}

	file, line := l.GetFileLine(2)
	l.postToRemote("INFO", "API Request info", file, r, responseTime, status, line)
}

func (l *Logger) postToRemote(level, msg, file string, r *http.Request, responseTime float64, status, line int) {
	client := &http.Client{}
	var method, req, ua, ip string
	if r != nil {
		method = r.Method
		req = r.RequestURI
		ua = r.UserAgent()
		ip = realip.RealIP(r)
	}

	reqData, err := json.Marshal(&Log{
		Level:        level,
		Timestamp:    time.Now().UTC(),
		Title:        msg,
		Message:      msg,
		AppName:      l.config.AppName,
		RefID:        l.ref,
		File:         file,
		Line:         strconv.Itoa(line),
		ResponseTime: responseTime,
		StatusCode:   status,
		Method:       method,
		Request:      req,
		UserAgent:    ua,
		IPAddress:    ip,
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
	if l.config.RemoteToken != "" {
		request.SetBasicAuth(l.config.RemoteUserName, l.config.RemoteToken)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("Unable to send log request. Err:", err)
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Println("Status code is not 200 (OK). Got:", response.StatusCode)
		// Handle the error codes
		return
	}

	return
}

func (l *Logger) PostToRemote(level, msg string) {
	file, line := l.GetFileLine(3)
	l.postToRemote(level, msg, file, nil, 0, 0, line)
}
