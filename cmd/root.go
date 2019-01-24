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

package cmd

import (
	"fmt"
	"os"

	"github.com/gozap/webmonitor/conf"
	"github.com/gozap/webmonitor/utils"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "webmonitor",
	Short: "A simple website monitor tool",
	Long: `
A simple website monitor tool.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./webmonitor.yaml)")
}

func initConfig() {
	if cfgFile == "" {
		cfgFile = "webmonitor.yaml"
		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			_, err = os.Create(cfgFile)
			utils.CheckAndExit(err)
			conf.Cfg = conf.Example()
			utils.CheckAndExit(conf.Cfg.WriteTo(cfgFile))
		} else if err != nil {
			utils.CheckAndExit(err)
		}
	}

	utils.CheckAndExit(conf.Cfg.LoadFrom(cfgFile))
}
