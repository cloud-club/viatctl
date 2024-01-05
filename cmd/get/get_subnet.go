/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	root "github.com/cloud-club/viatctl/cmd"

	"github.com/cloud-club/Aviator-service/pkg"
	"github.com/spf13/cobra"
)

// nodeCmd represents the create node command
var subnetCmd = &cobra.Command{
	Use:     "subnet",
	Aliases: []string{"subnets"},
	Short:   "Get subnet list",
	Long: `Get subnet list.
	
Example:
	viatctl get subnets
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp("subnet")

		response, err := ncp.Subnet.Get(pkg.VPC_API_URL + pkg.GET_SUBNET_LIST_PATH)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve subnet list.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(subnetCmd)
}
