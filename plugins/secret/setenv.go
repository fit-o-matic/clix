package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSetEnvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "setenv [key]",
		Short: "Set an environment variable with a secret value",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := NewStore()
			if err != nil {
				return fmt.Errorf("initializing store: %w", err)
			}
			if err := store.SetEnv(args[0]); err != nil {
				return fmt.Errorf("setting environment variable %q: %w", args[0], err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "environment variable %q set\n", args[0])
			return nil
		},
	}
}
