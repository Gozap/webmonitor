package monitor

import (
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/gozap/webmonitor/conf"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func Run() {
	conf.Cfg.Default()
	c := cron.New()

	for _, t := range conf.Cfg.Targets {
		tmpT := t
		logrus.Infof("monitoring target %s: %s", t.Name, t.Address)
		err := c.AddFunc(tmpT.Cron, func() {
			client := &http.Client{
				Timeout: tmpT.TimeOut,
			}

			if tmpT.Proxy != "" {
				p := func(_ *http.Request) (*url.URL, error) {
					return url.Parse(tmpT.Proxy)
				}
				client.Transport = &http.Transport{Proxy: p}
			}

			mon := HttpMonitor{
				client: client,
			}
			logrus.Infof("monitoring request: %s", tmpT.Address)
			err := mon.Monitor(tmpT)
			if err != nil {
				logrus.Error(err)
				// TODO: alarm
			}

		})
		if err != nil {
			logrus.Fatal(err)
		}
	}

	c.Start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}
