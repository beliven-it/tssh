/*
Copyright © 2023 Beliven
*/
package cmd

import (
	"fmt"
	"os"
	"tssh/defs"
	"tssh/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Version of the app provided
// in build phase
var Version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "tssh",
	Version: Version,
	Short:   "A CLI to easily sync, list, search and connect to Goteleport nodes",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/config.yml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if utils.Which("tsh") == "" {
		fmt.Println("Missing tsh executable")
		fmt.Println("Please follow the instructions for install the binaries here:")
		fmt.Print("\nhttps://goteleport.com/download/\n\n")
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".tssh" (without extension).
		viper.AddConfigPath(defs.ConfigFolderPath)
		viper.SetConfigType(defs.ConfigFileExtension)
		viper.SetConfigName(defs.ConfigFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match
}
