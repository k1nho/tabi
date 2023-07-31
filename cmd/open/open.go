// Package open contains the utilities for opening tabs
package open

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"

	"github.com/spf13/cobra"
)

type Options struct {
	// The file path to the file to be opened by tabi
	FilePath string
	// The URLs to be opened by tabi
	URLs []string
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
			opts.URLs = append(opts.URLs, args...)
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
		file, err := os.OpenFile(opts.FilePath, os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		err = json.NewDecoder(file).Decode(&urls)
		if err != nil {
			return err
		}

		bytes, err := os.ReadFile("series.json")
		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, &urls)
		if err != nil {
			return err
		}

		opts.URLs = append(opts.URLs, urls...)
	}

	switch osRuntime {
	case "windows":
		err = openTabWindows(opts.URLs)
	default:
		err = openTabUnix(opts.URLs)
	}

	return err
}

func openTabWindows(urls []string) error {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()
			cmd := exec.Command("/C", "start", "/B", "chrome", url)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Could not open %s: %v", url, err)
			}
		}(url, &wg)
	}
	wg.Wait()
	return nil
}

func openTabUnix(urls []string) error {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(url string, wg *sync.WaitGroup) {
			defer wg.Done()
			cmd := exec.Command("open", "-a", "Google Chrome", url)
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Could not open %s: %v", url, err)
			}
		}(url, &wg)
	}
	wg.Wait()
	return nil
}
