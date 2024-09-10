package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/supersonictw/socket_cat/cmd"
)

func main() {
	cobra.CheckErr(cmd.NewCLI().ExecuteContext(context.Background()))
}
