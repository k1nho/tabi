package update

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/k1nho/tab-cli/pkg/utils"
	"github.com/spf13/cobra"
)

const uptLongDesc string = `The upt command takes in a file path to a JSON file and updates the urls
// updating episode count by +1 of a url link in a JSON file
tabi series.json 
// updating episode count by -1 of a url link in a JSON file
tabi series.json -d
`

type Options struct {
	// Determine if the urls are updated up or down
	UpdateEpisodeDown bool

	// The file path containing the series to be updated
	FilePath string
}

func NewUpdateCommand() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "upt filepath [flags]",
		Short: "update urls on a given JSON file",
		Long:  uptLongDesc,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("Could not update series: file path must be provided")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.FilePath = args[0]
			return run(opts)

		},
	}

	cmd.Flags().BoolVarP(&opts.UpdateEpisodeDown, "down", "d", false, "update the episode count down")

	return cmd
}

func run(opts *Options) error {
	var series []utils.Series

	file, err := os.Open(opts.FilePath)
	if err != nil {
		return err
	}

	err = json.NewDecoder(file).Decode(&series)
	if err != nil {
		return err
	}

	// update
	if opts.UpdateEpisodeDown {
		for i := range series {
			if series[i].Ep-1 > 0 {
				series[i].UpdateDown()
			}
		}

		bytes, err := json.MarshalIndent(series, "", "\t")
		if err != nil {
			return err
		}

		err = os.WriteFile(opts.FilePath, bytes, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		for i := range series {
			if series[i].Ep+1 <= series[i].Limit {
				series[i].UpdateUp()
			}
		}

		bytes, err := json.MarshalIndent(series, "", "\t")
		if err != nil {
			return err
		}

		err = os.WriteFile(opts.FilePath, bytes, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
