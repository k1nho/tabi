package create

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

const longDescription = `Setup tabi workspace to manage the series files updates`

type opts struct {
	// Dir: the directory to setup tabi
	Dir string
}

func NewWorkspaceCommand() *cobra.Command {
	o := opts{Dir: ""}
	cmd := &cobra.Command{
		Use:   "workspace [flags]",
		Short: "Setup tabi",
		Long:  longDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run()
		},
	}

	cmd.Flags().StringVar(&o.Dir, "dir", "", "directory to setup tabi in")
	return cmd
}

func (o *opts) run() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	var tabiPath string
	if o.Dir != "" {
		tabiPath = path.Join(homeDir, o.Dir)
	} else {
		tabiPath = path.Join(homeDir, ".tabi")
	}

	_, err = os.Stat(tabiPath)
	// check that workspace exists
	if os.IsNotExist(err) {
		if err := os.MkdirAll(tabiPath, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	fmt.Println("Workspace for tabi has been setup!")
	return nil
}
