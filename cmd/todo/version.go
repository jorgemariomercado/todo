package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var version = "DEV"

func commandVersion() *cobra.Command {

	return &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf(
				"todo version: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n",
				version,
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH,
			)
		},
	}
}
