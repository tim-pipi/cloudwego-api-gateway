/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/fileutils"
)

var _ = path.Join

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates the template for an RPC server",
	Long: `Generates the default template for an RPC server to run on Kitex.
	This command will create a directory called kitex-template in the current directory,
	and will generate the kitex RPC code from the template.`,
	Run: func(cmd *cobra.Command, args []string) {
		idl := cmd.Flag("idl").Value.String()

		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		kitexDir := path.Join(dir, "kitex-template")
		os.Mkdir(kitexDir, 0777)

		fileutils.CopyTemplateDir(kitexDir)
		// Execute kitex command on the current directory
		kitexCmd := exec.Command(
			"kitex",
			"--thrift-plugin",
			"validator",
			"-module",
			"github.com/tim-pipi/cloudwego-api-gateway",
			"--template-dir",
			kitexDir,
			idl,
		)
		if out, err := kitexCmd.Output(); err != nil {
			fmt.Println(string(out))
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	// Add flag for IDL file
	genCmd.Flags().StringP("idl", "i", "", "Thrift IDL file to use for generating the RPC server")
	genCmd.MarkFlagRequired("idl")
}
