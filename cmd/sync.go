/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"3xd/internal"
	"3xd/internal/logger"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",
	Long: `Sync is a command which iterates over 100 int values and print them based on the provided flags.
		In default case it prints the message '100 lines was processed'. The import command has help, verbose and all flags.
		If we specify -h or --help flag it would print this magic message.
		With verbose command we are printing every 10 lines processed during the iteration of 100 int values. The message
		which is displayed is 'i%10 lines was processed'

		Flags:
			-h or --help => prints this magic message
			-v or --verbose => prints a log message after every 10 lines was processed
			-a or --all => add all message in the end of every log message`,
	ValidArgs: []string{internal.Help, internal.Verbose, internal.All},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("The sync process has been started")
		logMessage := " lines was processed"
		verbose, _ := cmd.Flags().GetBool(internal.Verbose)
		if allMsg, _  := cmd.Flags().GetBool(internal.All); allMsg {
			logMessage = fmt.Sprintf("%s\n%s", logMessage,internal.All)
		}
		logger.LogProcessedRows(verbose, logMessage)
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().BoolP(internal.Verbose, "v",false,
		"Some default value for -v or verbose")

	syncCmd.Flags().BoolP(internal.All, "a",false,
		"Some default value for -a or verbose")
}
