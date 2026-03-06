package clix

import (
	"github.com/spf13/cobra"
)

// Cli represents the main CLI application.
type Cli struct {
	Name        string
	Version     string
	Description string
	Cmd         *cobra.Command
}

// New creates a new Cli instance with the given name.
func New(name string) *Cli {
	return &Cli{
		Name: name,
		Cmd: &cobra.Command{
			Use: name,
		},
	}
}

// Execute runs the CLI application.
func (c *Cli) Execute() error {
	if c.Cmd.Short == "" && c.Description != "" {
		c.Cmd.Short = c.Description
	}
	if c.Cmd.Version == "" && c.Version != "" {
		c.Cmd.Version = c.Version
	}
	return c.Cmd.Execute()
}
