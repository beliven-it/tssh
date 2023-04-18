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

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:     "connect",
	Aliases: []string{"c"},
	Short:   "Connect through your goteleport remote nodes",
	Long: `This command allow you to search through a fuzzy search interface
the node you want connect.	
`,
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

		selection := connectionPresenter.Fzf(list)
		if selection == "" {
			fmt.Println("no connection selected")
			return
		}

		err = connectionService.Connect(selection)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
