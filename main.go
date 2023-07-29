package main

import (
	"log"

	"github.com/k1nho/tab-cli/cmd/root"
	"github.com/k1nho/tab-cli/pkg/utils"
)

func main() {
	rootCmd, err := root.NewRootCommand()
	if err != nil {
		log.Fatal(err)
	}

	utils.SetupRootCommand(rootCmd)

	err = rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
