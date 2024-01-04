/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package get

import (
	"fmt"

	root "github.com/cloud-club/viatctl/cmd"

	"github.com/cloud-club/Aviator-service/pkg"
	server "github.com/cloud-club/Aviator-service/types/server"
	"github.com/spf13/cobra"
)

// nodeCmd represents the create node command
var nodeCmd = &cobra.Command{
	Use:     "node",
	Aliases: []string{"nodes"},
	Short:   "Get nodes created",
	Long: `Get nodes created.
	
	Example:
		viatctl get nodes --region KR
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp()

		// Map command line arguments to request struct
		region, _ := cmd.Flags().GetString("region")

		lsr := &server.ListServerRequest{
			RegionCode: region,
		}

		response, err := ncp.Server.List(pkg.API_URL+pkg.GET_SERVER_INSTANCE_PATH, lsr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve nodes.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("region", "r", "", "Region code to get nodes")
}
