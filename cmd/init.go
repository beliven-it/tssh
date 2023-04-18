/*
Copyright © 2023 Beliven
*/
package cmd

import (
	"fmt"
	"tssh/defs"
	"tssh/interfaces"
	"tssh/services"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Initialize the configuration file and other assets",
	Long:    `Initialize the configuration file and other assets`,
	Run: func(cmd *cobra.Command, args []string) {
		configService := services.NewConfigService()

		err := configService.Init()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Initialize the configuration of ssh
		goteleport, err := interfaces.NewGoteleportInterface(
			viper.GetString(defs.ConfigKeyTeleportUser),
			viper.GetString(defs.ConfigKeyTeleportProxy),
			viper.GetBool(defs.ConfigKeyTeleportPasswordless),
		)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = goteleport.CreateSshConfig()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
