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

package conf

import (
	"errors"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

var Cfg Config

type Config struct {
	configPath string
	Basic      Basic   `yaml:"basic"`
	Targets    Targets `yaml:"targets"`
}

// set config file path
func (cfg *Config) SetConfigPath(configPath string) {
	cfg.configPath = configPath
}

// write config to yaml file
func (cfg Config) Write() error {
	if cfg.configPath == "" {
		return errors.New("config path not set")
	}
	out, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(cfg.configPath, out, 0644)
}

// load config from yaml file
func (cfg *Config) Load(filePath string) error {
	if filePath == "" {
		return errors.New("file path is empty")
	}
	cfg.configPath = filePath
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buf, cfg)
}

func (cfg *Config) Default() {

	var ts Targets

	for _, tmpT := range cfg.Targets {
		if tmpT.Cron == "" {
			tmpT.Cron = cfg.Basic.Cron
		}
		if tmpT.TimeOut == 0 {
			tmpT.TimeOut = cfg.Basic.TimeOut
		}
		if tmpT.Proxy == "" {
			tmpT.Proxy = cfg.Basic.Proxy
		}
		if tmpT.Method == "" {
			tmpT.Method = cfg.Basic.Method
		}
		ts = append(ts, tmpT)
	}
	cfg.Targets = ts
}

type Basic struct {
	TimeOut time.Duration `yaml:"timeout"`
	Method  string        `yaml:"method"`
	Proxy   string        `yaml:"proxy"`
	Cron    string        `yaml:"cron"`
}

type Target struct {
	Name    string        `yaml:"name"`
	Address string        `yaml:"address"`
	TimeOut time.Duration `yaml:"timeout"`
	Method  string        `yaml:"method"`
	Payload string        `yaml:"payload"`
	Proxy   string        `yaml:"proxy"`
	Cron    string        `yaml:"cron"`
}

type Targets []Target

func (ts Targets) Len() int {
	return len(ts)
}
func (ts Targets) Less(i, j int) bool {
	return ts[i].Name < ts[j].Name
}
func (ts Targets) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func Example() Config {
	return Config{
		Basic: Basic{
			Method:  "get",
			TimeOut: 5 * time.Second,
			Proxy:   "",
		},
		Targets: []Target{
			{
				Name:    "百度",
				Address: "https://baidu.com",
			},
			{
				Name:    "漠然",
				Address: "https://mritd.me",
			},
		},
	}
}
