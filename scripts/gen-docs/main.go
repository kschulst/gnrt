package main

import (
	"log"

	"github.com/spf13/cobra/doc"

	"github.com/kschulst/gnrt/cmd/gnrt/commands"
)

func main() {
	commands.RootCmd.DisableAutoGenTag = true

	if err := doc.GenMarkdownTree(commands.RootCmd, "./docs"); err != nil {
		log.Fatal(err)
	}
}
