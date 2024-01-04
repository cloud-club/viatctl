/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package stop

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
	Short: "Stop node",
	Long: `Stop node.
	
	Example:
		viatctl stop node --server serverNo

	Get server lists by running:
		viatctl get nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stoping node....")

		ncp := root.InitNcp()

		// Map command line arguments to request struct
		serverNo, _ := cmd.Flags().GetString("server")

		ssr := &server.StopServerRequest{
			ServerNo: serverNo,
		}

		response, err := ncp.Server.Stop(pkg.API_URL+pkg.STOP_SERVER_INSTANCE_PATH, ssr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to stop node.")
		}

		fmt.Println(response)
		fmt.Println("%s stopped successfully.", serverNo)
	},
}

func init() {
	stopCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("server", "s", "", "Server number to stop")
}
