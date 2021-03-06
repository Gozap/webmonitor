package main

import (
	"fmt"
	"runtime"

	"github.com/gozap/webmonitor/cmd"

	"github.com/spf13/cobra"
)

var versionTpl = `
Name: webmonitor
Version: %s
Arch: %s
BuildDate: %s
CommitID: %s
`

var (
	Version   string
	BuildDate string
	CommitID  string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long: `
Print version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(versionTpl, Version, runtime.GOOS+"/"+runtime.GOARCH, BuildDate, CommitID)
	},
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}
