/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/config"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the current services in the API Gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		config.CreateConfigDir()
		configPath := config.GetConfigPath()
		c, err := config.ReadConfig(configPath)

		if err != nil {
			return err
		}

		serviceMap, err := config.GetServiceMapFromDir(c.ThriftDir)

		if err != nil {
			return err
		}

		v, _ := cmd.Flags().GetBool("verbose")

		fmt.Printf("%-20s\n", "Service Name")
		for serviceName, filename := range serviceMap {
			if v {
				fmt.Printf("%-20s|\t%-20s\n", serviceName, filename)
			} else {
				fmt.Printf("%-20s\n", serviceName)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "Verbose output for the list command.")
}
