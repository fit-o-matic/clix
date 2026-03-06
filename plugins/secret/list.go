package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all secret keys",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			store, err := NewStore()
			if err != nil {
				return fmt.Errorf("initializing store: %w", err)
			}
			keys, err := store.List()
			if err != nil {
				return fmt.Errorf("listing secrets: %w", err)
			}
			for _, key := range keys {
				fmt.Fprintln(cmd.OutOrStdout(), key)
			}
			return nil
		},
	}
}
