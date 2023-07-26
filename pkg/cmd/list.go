/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/config"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/dir"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/service"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the current services in the API Gateway",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.ReadConfig()

		v, _ := cmd.Flags().GetBool("verbose")

		if ok, _ := dir.Exists(cfg.IDLDir); !ok {
			fmt.Println("IDL directory does not exist.\nPlease create the directory or set the env variable IDL_DIR to the correct path.")
			return nil
		}

		services, err := service.GetServicesFromIDLDir(cfg.IDLDir)

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
