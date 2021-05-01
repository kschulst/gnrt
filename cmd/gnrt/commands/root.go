package commands

import (
	"github.com/erdaltsksn/cui"
	"github.com/spf13/cobra"
)

type semver struct {
	Name        string
	Version     string
	Description string
}

// RootCmd represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "gnrt",
	Short: "Generate a data element",
	Long:  `Generate a data element useful for testing purposes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Success
		cui.Success(
			"Blah",
			"Foo bar",
		)
	},
}
