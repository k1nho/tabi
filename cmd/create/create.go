package create

import (
	"encoding/json"
	"errors"
	"fmt"
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

const createLongDesc = `THIS COMMAND IS USED TO CREATE URLS IN THE FORM (https://example.com/seriesname/ep)
The create command takes in a filepath to a JSON file with an array containing series in the form:
{
    url: string,
    ep: int,
    limit: int
}
then, creates a urls.json file which can be used by open to open all the urls
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

	cmd.Flags().StringVarP(&opts.Link, "add", "a", "", "A series link in the form (https://example.com/seriesname/ep) to be added to the JSON file")
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

	if opts.Link != "" {
		fmt.Println(opts.Link)
	}

	// create urls.json file
	urlsFile, err := os.Create("urls.json")
	if err != nil {
		return err
	}

	var urls []string
	for _, series := range series {
		url := fmt.Sprintf("%s/%d", series.BaseURL, series.Ep)
		urls = append(urls, url)
	}

	bytes, err = json.Marshal(urls)
	if err != nil {
		return err
	}

	_, err = urlsFile.Write(bytes)
	if err != nil {
		return err
	}

	return nil

}
