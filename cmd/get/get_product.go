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
var productCmd = &cobra.Command{
	Use:     "product",
	Aliases: []string{"products"},
	Short:   "Get product code list",
	Long: `Get product code list.
	
Example:
	viatctl get products --image-productcode SW.VSVR.OS.LNX64.CNTOS.0703.B050
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp("product")

		// Map command line arguments to request struct
		imageProductCode, _ := cmd.Flags().GetString("image-productcode")

		gpr := &server.GetProductRequest{
			ServerImageProductCode: imageProductCode,
		}

		response, err := ncp.ServerProduct.Get(pkg.API_URL+pkg.GET_SERVER_PRODUCT_LIST_PATH, gpr)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve product codes.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(productCmd)

	productCmd.Flags().StringP("image-productcode", "i", "", "Image product code to get product code list")
}
