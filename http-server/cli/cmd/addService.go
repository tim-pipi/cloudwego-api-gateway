/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/config"
)

// addServiceCmd represents the addService command
var addServiceCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new service to the API Gateway.",
	Long: `Add a new service to the API Gateway.
	Specify the idl file to create the new service using the --idl flag.
	Note that the idl file must be a valid Thrift file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("addService called")
		// Print the idl flag
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

		serviceName := services[0]
		err = c.Update(serviceName, idl)
		if err != nil {
			return fmt.Errorf("error updating configuration: %v", err)
		}

		err = config.CopyToConfigDir(idl, serviceName+".thrift")
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
	rootCmd.AddCommand(addServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addServiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addServiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Persistent flag for IDL file path
	addServiceCmd.PersistentFlags().StringP("idl", "i", "", "Path to IDL file")
	addServiceCmd.MarkPersistentFlagRequired("idl")
}
