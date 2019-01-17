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
		_ = c.AddFunc(tmpT.Cron, func() {

		})
	}
}
