package cmd

import (
	"fmt"
	"os"

	"github.com/gozap/webmonitor/monitor"

	"github.com/gozap/webmonitor/conf"
	"github.com/gozap/webmonitor/utils"
	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "webmonitor",
	Short: "A simple website monitor tool",
	Long: `
A simple website monitor tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		monitor.Run()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./webmonitor.yaml)")
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
