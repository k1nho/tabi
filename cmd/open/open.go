// Package open contains the utilities for opening tabs
package open

import (
	"encoding/json"
	"errors"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// Options are the options for the pizza bake command including user
// defined configurations
type Options struct {
	FilePath string
}

const openLongDesc string = `
The open command takes in a url/urls or a json file and opens all the urls on the web browser (chrome default)
    tabi open https://google.com
    tabi open https://google.com https://myanimelist.com
    tabi open --file=urls.json
`

// NewOpenCommand returns a new cobra command for 'tab open'
func NewOpenCommand() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "open url [flags]",
		Short: "open a given url/urls",
		Long:  openLongDesc,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 && opts.FilePath == "" {
				return errors.New("No url/urls provided to open")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.FilePath, "file", "f", "", "The JSON file to extract the URLS")

	return cmd
}

func run(opts *Options) error {
	osRuntime := runtime.GOOS

	var err error
	var urls []string

	// Open the json file if given
	if opts.FilePath != "" {
		bytes, err := os.ReadFile("series.json")

		err = json.Unmarshal(bytes, &urls)

		if err != nil {
			return err
		}
	}

	switch osRuntime {
	case "windows":
		err = openTabWindows()
	default:
		err = openTabUnix()
	}

	return err
}

func openTabWindows() error {
	return nil

}

func openTabUnix() error {
	return nil
}
