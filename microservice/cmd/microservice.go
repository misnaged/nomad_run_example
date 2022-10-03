package main

import (
	"microservice/cmd/root"
	"microservice/internal"
	"os"

	"microservice/cmd/serve"
)

func main() {
	app := &internal.App{}
	rootCmd := root.Cmd()
	rootCmd.AddCommand(serve.Cmd(app))

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
