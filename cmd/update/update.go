/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	"github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update node's SKU.",
	Long: `Update node's SKU.
	
	Get detail help in each subcommand
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Update called. No subcommands specified.")
	},
}

func init() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(updateCmd)
}
