/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/config"
)

// setIdlDirCmd represents the setIdlDir command
var setIdlDirCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the IDL file directory for the API Gateway.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("please provide the path to the IDL directory")
		}

		if len(args) > 1 {
			return fmt.Errorf("invalid number of arguments")
		}

		config.CreateConfigDir()
		configPath := config.GetConfigPath()

		c, err := config.ReadConfig(configPath)
		defer c.Write(configPath)

		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}

		if err := c.Update(args[0]); err != nil {
			return fmt.Errorf("error updating configuration: %v", err)
		}
		fmt.Println("IDL Directory updated successfully.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(setIdlDirCmd)
}
