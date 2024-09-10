package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:           "socket_cat",
		Short:         "Powerful relay server via WebSocket",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cmd.UsageString())
		},
	}

	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	rootCmd.AddCommand(
		newServerCLI(),
		newClientCLI(),
	)

	return rootCmd
}
