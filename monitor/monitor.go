/*
 * Copyright 2019 Gozap, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	t.CheckResponse(resp)

	return nil
}
