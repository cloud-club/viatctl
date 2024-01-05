/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package start

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
	Short: "Delete node for cluster",
	Long: `Delete node for cluster.

Example:
	viatctl start node --server serverNo

Get server lists by running:
	viatctl get nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting node....")

		ncp := root.InitNcp("server")

		// Map command line arguments to request struct
		serverNo, _ := cmd.Flags().GetString("server")

		ssr := &server.StartServerRequest{
			ServerNo: serverNo,
		}

		response, err := ncp.Server.Start(pkg.API_URL+pkg.START_SERVER_INSTANCE_PATH, ssr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to start node.")
		}

		fmt.Println(response)
		fmt.Println("%s started successfully.", serverNo)
	},
}

func init() {
	startCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("server", "s", "", "Server number to start")
}
