/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"fmt"

	root "github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create instance, services and etc.",
	Long: `Create instance, services and etc.
	
Get detail help in each subcommand.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create called. No subcommands specified.")
	},
}

func init() {
	rootCmd := root.GetRootCmd()
	rootCmd.AddCommand(createCmd)
}
