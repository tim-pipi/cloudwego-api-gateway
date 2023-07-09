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
		idlDir, err := config.GetDirFromConfig()

		if err != nil {
			return err
		}

		v, _ := cmd.Flags().GetBool("verbose")
		services, err := config.GetServicesFromIDLDir(idlDir)

		if err != nil {
			return err
		}

		fmt.Printf("%-20s\n", "Service Name")
		for _, service := range services {
			if v {
				fmt.Printf("%-20s|\t%-20s\n", service.Name, service.Path)

				for method, routes := range service.Routes {
					for _, route := range routes {
						fmt.Printf("%+20s|\t%-20s\n", method, route)
					}
				}
			} else {
				fmt.Printf("%-20s\n", service.Name)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("verbose", "v", false, "Verbose output for the list command.")
}
