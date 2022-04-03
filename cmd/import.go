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
	Short: "import iterates over 100 int values and print the values based on the flag type",
	Long: `
		Import is a command which iterates over 100 int values and print them based on the provided flags.
		In default case it prints the message '100 lines was processed'. The import command has help and verbose flag.
		If we specify -h or --help flag it would print this magic message.
		With verbose command we are printing every 10 lines processed during the iteration of 100 int values. The message
		which is displayed is 'i%10 lines was processed'

		Flags:
			-h or --help => prints this magic message
			-v or --verbose => prints a log message after every 10 lines was processed
	`,
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
