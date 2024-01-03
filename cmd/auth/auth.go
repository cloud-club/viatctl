package auth

import (
	"fmt"

	"github.com/cloud-club/viatctl/cmd"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication for viatctl",
	Long: `Authentication for viatctl to use Aviator-service API.

	Prerequisite: 
		Create API key and Secret key from NCP console. Look up(https://api.ncloud-docs.com/docs/common-ncpapi)

	You can get API key and Secret key from the console. Please keep the keys safe and do not share them with others.
	
	Example:
		viatctl auth --apikey APIKEY --secretkey SECRETKEY`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Authentication Called with")
		fmt.Println("API Key: ", cmd.Flag("apikey").Value)
		fmt.Println("Secret Key: ", cmd.Flag("secretkey").Value)

		// TODO: Implement authentication logic
		// 1. Initate service client with api & secret key
		// 2. Handle auth error
	},
}

func init() {
	rootCmd := cmd.GetRootCmd()
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().StringP("apikey", "a", "", "API key for authentication")
	authCmd.Flags().StringP("secretkey", "s", "", "Secret key for authentication")
}
