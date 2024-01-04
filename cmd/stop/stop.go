/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package stop

import (
	"fmt"

	"github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop instance, services and etc.",
	Long: `Stop instance, services and etc.
	
Get detail help in each subcommand
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stop called. No subcommands specified.")
	},
}

func init() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(stopCmd)
}
