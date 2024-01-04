/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/cloud-club/viatctl/cmd"
	_ "github.com/cloud-club/viatctl/cmd/auth"
	_ "github.com/cloud-club/viatctl/cmd/create"
	_ "github.com/cloud-club/viatctl/cmd/delete"
	_ "github.com/cloud-club/viatctl/cmd/get"
	_ "github.com/cloud-club/viatctl/cmd/start"
	_ "github.com/cloud-club/viatctl/cmd/stop"
	_ "github.com/cloud-club/viatctl/cmd/update"
)

func main() {
	cmd.Execute()
}
