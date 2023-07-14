/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/config"
	"github.com/tim-pipi/cloudwego-api-gateway/pkg/server"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the API Gateway HTTP server",
	Long: `Starts the API Gateway HTTP Server.

	Configuration is read from the environment variables.

	Environment variables:
	- $IDL_DIR is the directory containing the IDL files.
	- $ETCD_ADDR is the address of the etcd server.
	- $LOG_LEVEL is the log level.
	- $LOG_PATH is the path to the log file.
	- $ALLOW_METRICS is a flag to enable metrics.`,
	Run: func(cmd *cobra.Command, args []string) {
		svcConfig := config.ReadConfig()
		server.Start(svcConfig)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
