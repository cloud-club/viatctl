/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"

	"github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete instance, services and etc.",
	Long: `Delete instance, services and etc.
	
	Get detail help in each subcommand
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete called. No subcommands specified.")
	},
}

func init() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(deleteCmd)
}
