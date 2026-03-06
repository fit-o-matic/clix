package command

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/finkt/clix-kit/folder"
)

// RunContext provides the execution context for a command.
type RunContext struct {
	Args   []string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Env    map[string]string
}

type Descriptor struct {
	Summary       Summary `json:"summary"`
	ExecutionPath string  `json:"executionPath"`
}

type Summary struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Command struct {
	folder *folder.Folder
}

func Load(folder *folder.Folder) (*Command, error) {
	return &Command{folder: folder}, nil
}

func (c *Command) GetName() string {
	return c.folder.GetName()
}

func (c *Command) GetDescription() string {
	return "no description available"
}

// Run executes the command with the provided context.
// It looks for an executable with the same name as the command folder.
func (c *Command) Run(ctx *RunContext) error {
	execPath := filepath.Join(c.folder.GetPath(), c.folder.GetName())

	cmd := exec.Command(execPath, ctx.Args...)
	cmd.Stdin = ctx.Stdin
	cmd.Stdout = ctx.Stdout
	cmd.Stderr = ctx.Stderr

	// Build environment from current env plus context overrides
	cmd.Env = os.Environ()
	for k, v := range ctx.Env {
		cmd.Env = append(cmd.Env, k+"="+v)
	}

	return cmd.Run()
}
