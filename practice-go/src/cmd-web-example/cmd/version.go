package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version string
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show cmd-web version",
		Run: func(cmd *cobra.Command, args []string) {
			PrintVersion()
		},
	}
}

func getVersion() string {
	return version
}

func PrintVersion() {
	version := getVersion()
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Version", version)
}
