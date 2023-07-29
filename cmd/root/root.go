// Package root initiates and bootstraps the pizza CLI root Cobra command
package root

import (
	"github.com/k1nho/tab-cli/cmd/open"
	"github.com/spf13/cobra"
)

// NewRootCommand bootstraps a new root cobra command for the pizza CLI
func NewRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "tabi <command> <subcommand> [flags]",
		Short: "Tabi CLI",
		Long:  `A command line utility for opening and updating url links for series`,
		RunE:  run,
	}

	cmd.AddCommand(open.NewOpenCommand())

	return cmd, nil
}

func run(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}
