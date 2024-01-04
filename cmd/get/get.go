/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	"github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get instance, services and etc.",
	Long: `Get instance, services and etc.
	
Example:
	# To get server image product code list:
	viatctl get imageproducts
	
	# To get vpc list:
	viatctl get vpcs
	
	# To get subnet list:
	viatctl get subnets
	
	# To get access control group list:
	viatctl get accesscontrolgroups
	
	# To get server product code list:
	viatctl get products
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get called. No subcommands specified.")
	},
}

func init() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(getCmd)
}
