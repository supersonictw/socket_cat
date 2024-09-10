package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newServerCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "server",
		Short:         "Relay server",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cmd.UsageString())
		},
	}

	serverInitCmd := &cobra.Command{
		Use:   "init THE_SERVER_ENTRYPOINT_URL",
		Short: "Init server data",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		serverInitCmd,
	)

	return cmd
}

func serverRun(cmd *cobra.Command, args []string) {
	fmt.Println("Hello")
}
