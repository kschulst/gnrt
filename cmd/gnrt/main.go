package main

import (
	"github.com/erdaltsksn/cui"

	"github.com/kschulst/gnrt/cmd/gnrt/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		cui.Error("Something went wrong", err)
	}
}
