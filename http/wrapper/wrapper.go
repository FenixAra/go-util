package wrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/FenixAra/go-util/log"
)

const (
	JSON_DATA = iota
	FORM_DATA
)

//go:generate mockgen -source=wrapper.go -destination=mock_wrapper.go -package=wrapper
type IWrapper interface {
	Init(l *log.Logger)
	MakeRequest(method string, pType int, u string, payload interface{}, auth string, pass string, v interface{}) (int, interface{}, error)
}

type Wrapper struct {
	l *log.Logger
}

func New(l *log.Logger) *Wrapper {
	return &Wrapper{
		l: l,
	}
}

func (h *Wrapper) GetRequest(url string, v interface{}) (int, interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, nil, nil
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}

	err = json.Unmarshal(content, v)
	if err != nil {
		return 0, nil, err
	}
	return res.StatusCode, v, nil
}

func (h *Wrapper) MakeRequest(method string, pType int, u string, payload interface{}, auth string, pass string, v interface{}) (int, interface{}, error) {
	if auth == "" && method == "GET" {
		return h.GetRequest(u, v)
	}

	client := new(http.Client)
	var p []byte
	var err error
	switch pType {
	case JSON_DATA:
		p, err = json.Marshal(payload)
		if err != nil {
			h.l.Error("Unable to marshal Request:", u, ", Err:", err)
			return 0, nil, err
		}
	}

	req, err := http.NewRequest(method, u, bytes.NewBuffer(p))
	if err != nil {
		h.l.Error("Unable to create Request:", u, ", Err:", err)
		return 0, nil, err
	}

	if auth != "" {
		req.SetBasicAuth(auth, pass)
	}

	response, err := client.Do(req)
	if err != nil {
		h.l.Error("Unable to make Request:", u, ", Err:", err)
		return 0, nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return response.StatusCode, nil, nil
	}

	if v == nil {
		return response.StatusCode, nil, nil
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		h.l.Error("Unable to read response body, Err:", err)
		return 0, nil, err
	}

	err = json.Unmarshal(content, &v)
	if err != nil {
		h.l.Error("Unable to unmarshal block response, Err:", err)
		return 0, nil, err
	}

	return response.StatusCode, v, nil
}
