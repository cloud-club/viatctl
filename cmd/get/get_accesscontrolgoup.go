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
var acgCmd = &cobra.Command{
	Use:     "accesscontrolgroup",
	Aliases: []string{"accesscontrolgroups", "acg"},
	Short:   "Get access control group list",
	Long: `Get access control group list.

Example:
	viatctl get accesscontrolgroups
	viatctl get acg
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ncp := root.InitNcp()

		response, err := ncp.AccessControlGroup.Get(pkg.API_URL + pkg.GET_ACG_LIST_PATH)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to retrieve access control group list.")
		}

		fmt.Println(response)
	},
}

func init() {
	getCmd.AddCommand(acgCmd)
}
