package monitor

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gozap/webmonitor/conf"
)

type Monitor interface {
	Monitor(t conf.Target) error
}

type HttpMonitor struct {
	client *http.Client
}

func (m HttpMonitor) Monitor(t conf.Target) error {

	var req *http.Request
	var err error

	switch strings.ToUpper(t.Method) {
	case "GET":
		req, err = http.NewRequest("GET", t.Address, bytes.NewBufferString(t.Payload))
	case "POST":
		req, err = http.NewRequest("POST", t.Address, bytes.NewBufferString(t.Payload))
	}
	if err != nil {
		return err
	}

	resp, err := m.client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	err = t.CheckResponse(resp)
	if err != nil {
		return err
	}

	return nil
}
