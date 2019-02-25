package alarm

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gozap/webmonitor/conf"
)

type Alarm interface {
	Alarm(msg, level string) error
}

type Salicola struct {
	conf.Salicola
}

func (sa Salicola) Alarm(msg, level string) error {

	form := url.Values{}
	form.Set("subject", "Web Monitor")
	form.Set("level", level)
	form.Set("message", fmt.Sprintf(msg))

	req, err := http.NewRequest("POST", sa.Address, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+sa.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, resp.Body)
		return errors.New(fmt.Sprintf("alarm send failed, response code %d: %s", resp.StatusCode, buf.String()))
	}
	return nil

}
