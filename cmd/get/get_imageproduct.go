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
var ipcCmd = &cobra.Command{
	Use:     "image-productcode",
	Aliases: []string{"image-productcodes", "ipc"},
	Short:   "Get image product code list",
	Long: `Get image product code list.
	
	Example:
		viatctl get image-productcodes
		viatctl get ipc
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp()

		response, err := ncp.ServerImageProduct.Get(pkg.API_URL + pkg.GET_SERVER_IMAGE_PRODUCT_LIST_PATH)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve image product code list.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(ipcCmd)
}
