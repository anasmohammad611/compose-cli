/*
Copyright © 2023 Shuvojit Sarkar <shuvojit@healtih.in>
*/
package cmd

import (
	"github.com/anasmohammad611/compose-cli/common/printer"
	"github.com/spf13/cobra"
)

func initRun(cmd *cobra.Command, args []string) {
	printer.Info("init called")
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: initRun,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
