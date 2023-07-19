/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/config"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/dir"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/fileutils"
	"github.com/tim-pipi/cloudwego-api-gateway/internal/service"
)

var _ = path.Join

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates the template for an RPC server",
	Long: `Generates the default template for an RPC server to run on Kitex.
This command will create a directory called kitex-template in the current directory,
and will generate the kitex RPC code from the template.`,
	Example: `cwgo gen -i path/to/idl/file.thrift Generate RPC server code from the specified IDL file
cwgo gen -s HelloService Generate RPC server code for the specified service
	`,
	Run: func(cmd *cobra.Command, args []string) {
		idl := cmd.Flag("idl").Value.String()
		serviceName := cmd.Flag("service").Value.String()

		if (serviceName == "") && (idl == "") {
			fmt.Println("Please specify either an IDL file or a service name.")
			os.Exit(1)
		}

		if serviceName != "" {
			cfg := config.ReadConfig()

			if ok, _ := dir.Exists(cfg.IDLDir); !ok {
				fmt.Println("IDL directory does not exist.\nPlease create the directory or set the env variable IDL_DIR to the correct path.")
				return
			}

			services, err := service.GetServicesFromIDLDir(cfg.IDLDir)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			found := false
			for _, service := range services {
				if strings.EqualFold(service.Name, serviceName) {
					idl = service.Path
					found = true
					break
				}
			}

			if !found {
				fmt.Println("No service with that name found.")
				os.Exit(1)
			}
		}

		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		kitexDir := path.Join(dir, "kitex-template")
		os.Mkdir(kitexDir, 0777)

		fileutils.CopyTemplateKitexDir(kitexDir)
		fileutils.CopyTemplateFile("Dockerfile", path.Join(dir, "Dockerfile"))
		module := cmd.Flag("module").Value.String()
		// Execute kitex command on the current directory
		kitexCmd := exec.Command(
			"kitex",
			"--thrift-plugin",
			"validator",
			"-module",
			module,
			"--template-dir",
			kitexDir,
			idl,
		)

		kitexCmd.Stdout = os.Stdout
		kitexCmd.Stderr = os.Stderr

		if err := kitexCmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	// Add flag for IDL file
	genCmd.Flags().StringP("idl", "i", "", "Thrift IDL file to use for generating RPC server code")
	genCmd.Flags().StringP("service", "s", "", "Service name to use for generating RPC server code")
	genCmd.MarkFlagsMutuallyExclusive("idl", "service")

	genCmd.Flags().StringP("module", "m", "", "Go module name")
	genCmd.MarkFlagRequired("module")
}
