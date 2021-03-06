package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

var (
	verbose bool
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "", false,
		"Print more information to STDOUT")
}

// RootCmd represents the root command of this tool
var RootCmd = &cobra.Command{
	Use:   "ecs-bind",
	Short: "Executes a command with ecs dynamic meta value as environment",
	Args: func(cmd *cobra.Command, args []string) error {
		dashIx := cmd.ArgsLenAtDash()
		if dashIx == -1 {
			return errors.New("please separate services and command with '--'. See usage")
		}
		return nil
	},
	RunE: execRun,
}

// Execute is the function to run root command of this tool
func Execute() {
	if _, err := RootCmd.ExecuteC(); err != nil {
		os.Exit(1)
	}
}
