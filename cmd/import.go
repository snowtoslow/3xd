/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"3xd/internal"
	"3xd/internal/logger"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: []string{internal.Verbose, internal.Help},
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := cmd.Flags().GetBool(internal.Verbose)
		logger.LogProcessedRows(verbose, " lines was processed")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.Flags().BoolP(internal.Verbose, "v", false,
		"Some default value for -v or verbose")
}
