package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// Version 版本號
var Version string

const (
	// App 程式名
	App = "revel-i18n"
)

var v bool
var rootCmd = &cobra.Command{
	Use:   App,
	Short: "revel i18n tools",
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			fmt.Println(Version)
		} else {
			fmt.Println(App)
			fmt.Println(Version)
			fmt.Printf(`Use "%v --help" for more information about this program.
`, App)
		}
	},
}

func init() {
	flags := rootCmd.Flags()
	flags.BoolVarP(&v,
		"version",
		"v",
		false,
		"show version",
	)
}

// Execute 執行命令
func Execute() error {
	return rootCmd.Execute()
}
func abort() {
	os.Exit(1)
}
