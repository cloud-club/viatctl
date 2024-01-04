/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

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
	Short: "Create node for cluster",
	Long: `Create node for cluster.

Example:
	viatctl create node --image-productcode SW.VSVR.OS.LNX64.CNTOS.0703.B050 --vpc-no 52833 --subnet-no 120320 --network-interface-order 0 --access-control-group 148207 --productcode SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002

1. To get server image product code list:
	viatctl get imageproducts
2. To get vpc list:
	viatctl get vpcs
3. To get subnet list:
	viatctl get subnets
4. To get access control group list:
	viatctl get accesscontrolgroups
5. To get server product code list:
	viatctl get products
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating node....")

		ncp := root.InitNcp()

		// Map command line arguments to request struct
		serverImageProductCode, _ := cmd.Flags().GetString("image-productcode")
		vpcNo, _ := cmd.Flags().GetString("vpc")
		subnetNo, _ := cmd.Flags().GetString("subnet")
		networkInterfaceOrder, _ := cmd.Flags().GetInt("network-interface-order")
		accessControlGroupNo, _ := cmd.Flags().GetString("access-control-group")
		serverProductCode, _ := cmd.Flags().GetString("productcode")

		csr := &server.CreateServerRequest{
			ServerImageProductCode:    serverImageProductCode,
			VpcNo:                     vpcNo,
			SubnetNo:                  subnetNo,
			NetworkInterfaceOrder:     networkInterfaceOrder,
			AccessControlGroupNoListN: accessControlGroupNo,
			ServerProductCode:         serverProductCode}

		response, err := ncp.Server.Create(pkg.API_URL+pkg.CREATE_SERVER_INSTANCE_PATH, csr, []int{1, 1})
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to create node.")
		}

		fmt.Println(response)
		fmt.Println("Node created successfully.")
	},
}

func init() {
	createCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("image-productcode", "i", "", "Server image product code")
	nodeCmd.Flags().StringP("vpc", "v", "", "VPC number")
	nodeCmd.Flags().StringP("subnet", "s", "", "Subnet number")
	nodeCmd.Flags().IntP("network-interface-order", "n", 0, "Network interface order")
	nodeCmd.Flags().StringP("access-control-group", "a", "", "Access control group number")
	nodeCmd.Flags().StringP("productcode", "p", "", "Server product code")
}
