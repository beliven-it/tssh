/*
Copyright © 2023 Beliven
*/
package cmd

import (
	"fmt"
	"os"
	"tssh/defs"
	"tssh/interfaces"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var logoutCmd = &cobra.Command{
	Use:     "logout",
	Aliases: []string{"d"},
	Short:   "Logout from all clusters",
	Long:    `This command allow you to logout from the cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		goteleportService, err := interfaces.NewGoteleportInterface(
			viper.GetString(defs.ConfigKeyTeleportUser),
			viper.GetString(defs.ConfigKeyTeleportProxy),
			viper.GetBool(defs.ConfigKeyTeleportPasswordless),
		)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		goteleportService.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
