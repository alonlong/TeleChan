package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// the application's config file
var configFile string

// RootCmd represents the base command when called without any subcommand
var RootCmd = &cobra.Command{
	Use:   "TeleChan",
	Short: "The TeleChan server application",
	Long:  "Run the 'serve' subcommand to start the grpc server",
}

// Execute adds all child command to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&configFile, "config", "./config.yml", "server config file")
}
