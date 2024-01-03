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
		aviator create node --serverimage-productcode SW.VSVR.OS.LNX64.CNTOS.0703.B050 --vpc-no 52833 --subnet-no 120320 --network-interface-order 0 --access-control-group-no 148207 --server-productcode SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002

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
		fmt.Println("create node called")

		ncp := root.InitNcp()

		// Map command line arguments to request struct
		serverImageProductCode, _ := cmd.Flags().GetString("serverimage-productcode")
		vpcNo, _ := cmd.Flags().GetString("vpc-no")
		subnetNo, _ := cmd.Flags().GetString("subnet-no")
		networkInterfaceOrder, _ := cmd.Flags().GetInt("network-interface-order")
		accessControlGroupNo, _ := cmd.Flags().GetString("access-control-group-no")
		serverProductCode, _ := cmd.Flags().GetString("server-productcode")

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
		}

		fmt.Println(response)
	},
}

func init() {
	createCmd.AddCommand(nodeCmd)

	nodeCmd.Flags().StringP("serverimage-productcode", "i", "", "Server image product code")
	nodeCmd.Flags().StringP("vpc-no", "v", "", "VPC number")
	nodeCmd.Flags().StringP("subnet-no", "s", "", "Subnet number")
	nodeCmd.Flags().IntP("network-interface-order", "n", 0, "Network interface order")
	nodeCmd.Flags().StringP("access-control-group-no", "a", "", "Access control group number")
	nodeCmd.Flags().StringP("server-productcode", "p", "", "Server product code")
}
