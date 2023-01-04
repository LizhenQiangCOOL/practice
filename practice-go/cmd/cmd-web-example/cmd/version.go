package cmd

import (
	"my-web/internal/utils"

	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show cmd-web version",
		Run: func(cmd *cobra.Command, args []string) {
			utils.PrintVersion()
		},
	}
}
