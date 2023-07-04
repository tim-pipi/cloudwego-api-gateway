/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/config"
)

// updateServiceCmd represents the update command
var updateServiceCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an existing service in the API Gateway",
	Long: `Updates an existing service in the API Gateway.
	Specify the idl file to update using the --idl flag.
	Note that the idl file must be a valid Thrift file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		idl, _ := cmd.Flags().GetString("idl")

		services, err := config.GetServicesFromIDL(idl)

		if err != nil {
			return err
		}

		err = CheckServices(services)
		if err != nil {
			return err
		}

		config.CreateConfigDir()
		configPath := config.GetConfigPath()
		c, err := config.ReadConfig(configPath)
		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}
		configServiceNames := c.ServiceNameToThriftFile
		_, ok := configServiceNames[services[0]]
		if !ok {
			return fmt.Errorf("service does not exist")
		}

		err = config.CopyToConfigDir(idl, services[0]+".thrift")
		if err != nil {
			return fmt.Errorf("error copying IDL file to config directory: %v", err)
		}
		
		err = c.Write(configPath)
		if err != nil {
			return fmt.Errorf("error writing configuration: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateServiceCmd)

	updateServiceCmd.PersistentFlags().StringP("idl", "i", "", "Path to IDL file")
	updateServiceCmd.MarkPersistentFlagRequired("idl")
}
