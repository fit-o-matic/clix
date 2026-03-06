package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set [key] [value]",
		Short: "Set a secret",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := NewStore()
			if err != nil {
				return fmt.Errorf("initializing store: %w", err)
			}
			if err := store.Set(args[0], args[1]); err != nil {
				return fmt.Errorf("setting secret %q: %w", args[0], err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), "secret %q saved\n", args[0])
			return nil
		},
	}
}
