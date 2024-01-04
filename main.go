/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/cloud-club/viatctl/cmd"
)

func main() {
	os.Args = []string{"", "create"} // Simulate running "viatctl auth"
	cmd.Execute()
}
