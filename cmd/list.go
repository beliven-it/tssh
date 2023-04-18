/*
Copyright © 2023 Beliven
*/
package cmd

import (
	"fmt"
	"os"
	"tssh/defs"
	"tssh/presenters"
	"tssh/services"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List your remote nodes",
	Long:    `This command allow you to list all nodes you can access`,
	Run: func(cmd *cobra.Command, args []string) {

		connectionService, err := services.NewConnectionService(
			viper.GetString(defs.ConfigKeyTeleportUser),
			viper.GetString(defs.ConfigKeyTeleportProxy),
			viper.GetBool(defs.ConfigKeyTeleportPasswordless),
		)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		connectionPresenter := presenters.NewConnectionPresenter()

		list, err := connectionService.ListConnections()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Print(
			connectionPresenter.Text(list),
		)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
