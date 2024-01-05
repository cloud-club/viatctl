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
var niCmd = &cobra.Command{
	Use:     "networkinterface",
	Aliases: []string{"networkinterfaces", "ni"},
	Short:   "Get network interface list",
	Long: `Get network interface list.
	
Example:
	viatctl get networkinterfaces
	viatctl get ni
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp("networkinterface")

		response, err := ncp.Network.Get(pkg.API_URL + pkg.GET_NETWORKINTERFACE_LIST_PATH)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve network interface list.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(niCmd)
}
