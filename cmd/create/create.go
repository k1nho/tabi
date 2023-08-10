package create

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/k1nho/tab-cli/pkg/utils"
	"github.com/spf13/cobra"
)

type Options struct {
	// The file path to the JSON file to be opened by tabi
	FilePath string
	// Optional link given to be added to the JSON file
	Link string
}

const createLongDesc = ` THIS COMMAND IS USED FOR URLS IN THE FORM (https://example.com/seriesname/ep)
The create command takes in a filepath to a JSON file with an array containing series in the form:
    {
        url: string,
        ep: int,
        limit: int
    }
and creates a urls.json file which can be used to open all the urls
`

func NewCreateCommand() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "create filepath [flags]",
		Short: "create a JSON file containing all the URLS to be opened",
		Long:  createLongDesc,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("No file path provided")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.FilePath = args[0]
			return run(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Link, "", "a", "", "A series link in the form (https://example.com/seriesname/ep) to be added to the JSON file")
	return cmd
}

func run(opts *Options) error {
	var series []utils.Series

	// READ FILE
	bytes, err := os.ReadFile(opts.FilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &series)
	if err != nil {
		return err
	}

	fmt.Printf("the series are %v", series)

	if opts.Link != "" {
		url, err := url.Parse(opts.Link)
		if err != nil {
			return err
		}
		fmt.Println(url.Path)
	}

	return nil

}
