/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tim-pipi/cloudwego-api-gateway/http-server/internal/pkg/config"
)

// generateRpcCmd represents the generateRpc command
var generateRpcCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates the scaffolding code for the RPC Server.",
	Long: `Generates the scaffolding code for the RPC Server
	based on a specific Thrift file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		idlDir, err := config.GetDirFromConfig()
		if err != nil {
			return err
		}

		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		goCmd := exec.Command("go", "list", "-m")
		goCmd.Dir = dir

		output, err := goCmd.Output()
		if err != nil {
			return err
		}

		moduleName := strings.TrimSpace(string(output))
		kitexStr := fmt.Sprintf("kitex -module %s %s", moduleName, idlDir)

		fmt.Printf("Running kitex command:\n%s\n", kitexStr)

		kitexArgs := strings.Split(kitexStr, " ")
		kitexCmd := exec.Command(kitexArgs[0], kitexArgs[1:]...)

		pipereader, pipewriter := io.Pipe()
		kitexCmd.Stdout = pipewriter
		kitexCmd.Stderr = pipewriter

		go func() {
			defer pipewriter.Close()
			scanner := bufio.NewScanner(pipereader)

			for scanner.Scan() {
				fmt.Println("kitex: ", scanner.Text())
			}
		}()

		// kitexCmd.Stdout = os.Stdout
		// kitexCmd.Stderr = os.Stderr

		err = kitexCmd.Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateRpcCmd)
}
