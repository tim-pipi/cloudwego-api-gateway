/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// generateRpcCmd represents the generateRpc command
var generateRpcCmd = &cobra.Command{
	Use:   "generateRpc",
	Short: "Generates the scaffolding code for the RPC server.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sampleCommand := `kitex -module github.com/tim-pipi/cloudwego-api-gateway/http-server ../../idl/hello_api.thrift`

		// Execute the sample command
		kitexArgs := strings.Split(sampleCommand, " ")
		exec.Command(kitexArgs[0], kitexArgs[1:]...).Run()

		fmt.Println("generateRpc called")
	},
}

func init() {
	rootCmd.AddCommand(generateRpcCmd)
}
