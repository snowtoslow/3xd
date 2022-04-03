/*
Copyright © 2022 Vladimir Leadavschi snowtoslow snowtoslow@gmail.com

*/
package cmd

import (
	"3xd/config"
	"3xd/internal"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "3xd",
	Short: "3xd small cli application",
	Long: `3xd is a CLI	 application with functionality of printing logs during the process of iterating over 100 number.
		Created using cobra-cli tool
		It ha 2 base commands:
			1. import 
			2. sync
			3. config - we can specify a config file in json or yml format, which would be parsed using viper utility.
		also it contains help command which will display this magic msg and will show us the flags of 3xd.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString(internal.Config)
		cfg, err := config.ParseConfigFile(configPath)
		if err != nil{
			log.Println("Failed to parse config file: ", err)
		}

		log.Printf("Cfg content: %+v", cfg)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config", "config file (default is $HOME/.3xd.yaml)")
}


