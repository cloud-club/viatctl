/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"

	root "github.com/cloud-club/viatctl/cmd"

	"github.com/cloud-club/Aviator-service/pkg"
	server "github.com/cloud-club/Aviator-service/types/server"
	"github.com/spf13/cobra"
)

// nodeCmd represents the create node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Update node",
	Long: `Update node. Requires server number and server product code that you want to update to. 
	
Example:
	viatctl update node --server serverNo --productcode SVR.VSVR.STAND.C032.M128.NET.HDD.B050.G002

Get server lists by running:
	viatctl get nodes

Get server product code list by running:
	viatctl get productcodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Updating node....")

		ncp := root.InitNcp("server")

		// Map command line arguments to request struct
		serverNo, _ := cmd.Flags().GetString("server")
		productCode, _ := cmd.Flags().GetString("productcode")

		usr := &server.UpdateServerRequest{
			ServerInstanceNo:  serverNo,
			ServerProductCode: productCode,
		}

		response, err := ncp.Server.Update(pkg.API_URL+pkg.UPDATE_SERVER_INSTANCE_PATH, usr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to update node.")
		}

		fmt.Println(response)
		fmt.Println("%s updated successfully to %s.", serverNo, productCode)
	},
}

func init() {
	updateCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("server", "s", "", "Server number to update")
	nodeCmd.Flags().StringP("productcode", "p", "", "Server product code to update to")
}
