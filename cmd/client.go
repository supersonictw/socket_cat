package cmd

import "github.com/spf13/cobra"

func newClientCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "client",
		Short:         "An example relay client reads messages from stdin",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Print(cmd.UsageString())
		},
	}

	connectCmd := &cobra.Command{
		Use:   "connect",
		Short: "An example relay client reads messages from stdin",
		Args:  cobra.ExactArgs(1),
	}
	connectCmd.Flags().BoolP("no-best-latency", "fn", false, "Name of the Modelfile")

	return cmd
}

func clientRun(cmd *cobra.Command, args []string) {

}
