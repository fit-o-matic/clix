// Package main implements the secret plugin for clix.
package main

import (
	kit "github.com/finkt/clix-kit"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "manages secrets for clix",
	}
	cmd.AddCommand(
		newSetCmd(),
		newListCmd(),
		newDeleteCmd(),
		newSetEnvCmd(),
	)
	return cmd
}

func main() {
	p := &kit.Plugin{
		Name:        "secret",
		Description: "manages secrets for clix",
		Version:     "0.1.0",
		Usage:       "secret [get|set|list]",
		Cmd:         newRootCmd(),
	}
	p.Execute()
}
