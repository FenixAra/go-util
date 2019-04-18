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
	Level          string    `json:"level"`
	Timestamp      time.Time `json:"timestamp"`
	Title          string    `json:"message"`
	Message        string    `json:"full_message"`
	AppName        string    `json:"app_name"`
	RefID          string    `json:"ref_id"`
	File           string    `json:"file"`
	Line           string    `json:"line"`
	ResponseTime   float64   `json:"response_time"`
	StatusCode     int       `json:"status_code"`
	Method         string    `json:"method"`
	Request        string    `json:"request"`
	UserAgent      string    `json:"user_agent"`
	CustomerID     string    `json:"customer_id"`
	IPAddress      string    `json:"ip_address"`
	RequestGroup   string    `json:"request_group" example:"Ping"`
	AppVersion     string    `json:"app_version" example:"App Version"`
	TimeTaken      float64   `json:"time_taken" example:"1.11"`
	DependancyType string    `json:"dependancy_type" example:"http,database"`
	DependancyName string    `json:"dependancy_name" example:"googleapi,booktripsp"`
}

var logChan = make(chan Log, 2000)

func (l *Logger) Log(lg *Log) {
	lg.Level = Info
	lg.Timestamp = time.Now().UTC()
	lg.RefID = l.config.Reference
	lg.AppName = l.config.AppName
	file, line := l.GetFileLine(2)
	lg.File = file
	lg.Line = strconv.Itoa(line)
	logChan <- *lg
}

func (l *Logger) LogAPIInfo(r *http.Request, responseTime float64, status int) {
	if !l.config.remoteLogger {
		return
	}

	file, line := l.GetFileLine(2)
	l.postToRemote("INFO", "API Request info", file, r, responseTime, status, line)
}

func (l *Logger) postToRemote(level, msg, file string, r *http.Request, responseTime float64, status, line int) {
	var method, req, ua, ip string
	if r != nil {
		method = r.Method
		req = r.RequestURI
		ua = r.UserAgent()
		ip = realip.RealIP(r)
	}
	logData := &Log{
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
		UserAgent:    ua,
		IPAddress:    ip,
	}

	logChan <- *logData

}

func (l *Logger) SendLogs() {
	for {
		logData := <-logChan
		client := &http.Client{}
		reqData, err := json.Marshal(logData)
		if err != nil {
			log.Println("Unable to marshal log request. Err:", err)
			break
		}

		request, err := http.NewRequest(http.MethodPost, l.config.RemoteLoggerURL, bytes.NewBuffer(reqData))
		if err != nil {
			log.Println("Unable to create new log request. Err:", err)
			break
		}

		// Setting all headers
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		if l.config.RemoteToken != "" {
			request.SetBasicAuth(l.config.RemoteUserName, l.config.RemoteToken)
		}

		response, err := client.Do(request)
		if err != nil {
			log.Println("Unable to send log request. Err:", err)
			break
		}

		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			log.Println("Status code is not 200 (OK). Got:", response.StatusCode)
			// Handle the error codes
			break
		}
	}
}

func (l *Logger) PostToRemote(level, msg string) {
	file, line := l.GetFileLine(3)
	l.postToRemote(level, msg, file, nil, 0, 0, line)
}
