/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package delete

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
	viatctl delete node --server serverNo

Get server lists by running:
	viatctl get nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating node....")

		ncp := root.InitNcp("server")

		// Map command line arguments to request struct
		serverNo, _ := cmd.Flags().GetString("server")

		dsr := &server.DeleteServerRequest{
			ServerNo: serverNo,
		}

		response, err := ncp.Server.Delete(pkg.API_URL+pkg.DELETE_SERVER_INSTANCE_PATH, dsr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to delete node.")
		}

		fmt.Println(response)
		fmt.Println("%s deleted successfully.", serverNo)
	},
}

func init() {
	deleteCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("server", "s", "", "Server number to delete")
}
