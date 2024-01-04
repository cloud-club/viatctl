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
var vpcCmd = &cobra.Command{
	Use:     "vpc",
	Aliases: []string{"vpcs"},
	Short:   "Get vpc list",
	Long: `Get vpc list.
	
	Example:
		viatctl get vpcs
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp()

		response, err := ncp.Vpc.Get(pkg.API_URL + pkg.GET_VPC_LIST)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve vpc list.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(vpcCmd)
}
