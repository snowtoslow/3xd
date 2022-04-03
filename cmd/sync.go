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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: []string{internal.Verbose, internal.All, internal.Help},
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
