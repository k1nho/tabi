package create

import (
	"github.com/spf13/cobra"
)

const createLongDesc = "the create command generates the tabi workspace and the series file"

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create filepath [flags]",
		Short: "create a JSON file containing all the URLS to be opened",
		Long:  createLongDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(NewFileCommand())
	cmd.AddCommand(NewWorkspaceCommand())
	return cmd
}
