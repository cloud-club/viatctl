/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"fmt"

	"github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start instance, services and etc.",
	Long: `Start instance, services and etc.
	
	Get detail help in each subcommand
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start called. No subcommands specified.")
	},
}

func init() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(startCmd)
}
