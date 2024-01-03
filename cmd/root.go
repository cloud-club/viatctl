/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	pkg "github.com/cloud-club/Aviator-service/pkg"
	"github.com/spf13/cobra"
)

func InitNcp() *pkg.NcpService {
	ncp := pkg.NewNcpService("token")
	ncp.Server = pkg.NewServerService(os.Getenv("API_KEY"), os.Getenv("SECRET_KEY"))

	return ncp
}

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "viatctl",
	Short: "CLI for Aviator",
	Long: `CLI for Aviator.

Aviator is a tool for OS provisioning using custom resource definitions (CRDs) and operators for vanilla k8s deployed in Naver Cloud Platform.

Examples:
  # Define API key and Secret key for NCP
  viatctl auth --apikey APIKEY --secretkey SECRETKEY

  # Create new node(VM)
  viatctl create node --vpc vpcNo --subnet subnetNo --imageproduct imageProductCode --product productCode

  # Stop node
  viatctl stop node --server serverNo

  # Get nodes
  viatctl get nodes

  # Update node
  viatctl update node --server serverNo --product productCode

  # Delete node
  viatctl delete node --server serverNo

  # Get image product list
  viatctl get imageproducts

  # Get product list
  viatctl get products

  # Get vpc list
  viatctl get vpcs

  # Get subnet list
  viatctl get subnets

  # Get accesscontrolgroup list
  viatctl get accesscontrolgroups

  # Get networkinterface list
  viatctl get networkinterfaces`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func GetRootCmd() *cobra.Command {
	return RootCmd
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.viatctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
