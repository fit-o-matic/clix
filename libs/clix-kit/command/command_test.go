package command

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/finkt/clix-kit/folder"
)

func TestLoad(t *testing.T) {
	tmpDir := t.TempDir()
	f := folder.New(tmpDir)

	cmd, err := Load(f)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmd == nil {
		t.Fatal("expected command to be non-nil")
	}
}

func TestGetName(t *testing.T) {
	tmpDir := t.TempDir()
	cmdDir := filepath.Join(tmpDir, "mycommand")
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}

	cmd, err := Load(folder.New(cmdDir))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cmd.GetName() != "mycommand" {
		t.Errorf("expected name %q, got %q", "mycommand", cmd.GetName())
	}
}

func TestGetDescription(t *testing.T) {
	tmpDir := t.TempDir()
	cmd, err := Load(folder.New(tmpDir))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	desc := cmd.GetDescription()
	if desc == "" {
		t.Error("expected non-empty description")
	}
}

func TestRun_Success(t *testing.T) {
	tmpDir := t.TempDir()
	cmdName := "testcmd"
	cmdDir := filepath.Join(tmpDir, cmdName)
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}

	// Create a simple shell script that echoes args
	script := `#!/bin/sh
echo "args: $@"
`
	scriptPath := filepath.Join(cmdDir, cmdName)
	if err := os.WriteFile(scriptPath, []byte(script), 0755); err != nil {
		t.Fatalf("failed to write script: %v", err)
	}

	cmd, err := Load(folder.New(cmdDir))
	if err != nil {
		t.Fatalf("failed to load command: %v", err)
	}

	var stdout bytes.Buffer
	ctx := &RunContext{
		Args:   []string{"arg1", "arg2"},
		Stdin:  strings.NewReader(""),
		Stdout: &stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if !strings.Contains(output, "arg1") || !strings.Contains(output, "arg2") {
		t.Errorf("expected output to contain args, got: %q", output)
	}
}

func TestRun_WithEnv(t *testing.T) {
	tmpDir := t.TempDir()
	cmdName := "testcmd"
	cmdDir := filepath.Join(tmpDir, cmdName)
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}

	// Create a script that prints an env var
	script := `#!/bin/sh
echo "MY_VAR=$MY_VAR"
`
	scriptPath := filepath.Join(cmdDir, cmdName)
	if err := os.WriteFile(scriptPath, []byte(script), 0755); err != nil {
		t.Fatalf("failed to write script: %v", err)
	}

	cmd, err := Load(folder.New(cmdDir))
	if err != nil {
		t.Fatalf("failed to load command: %v", err)
	}

	var stdout bytes.Buffer
	ctx := &RunContext{
		Args:   []string{},
		Stdin:  strings.NewReader(""),
		Stdout: &stdout,
		Stderr: os.Stderr,
		Env:    map[string]string{"MY_VAR": "hello"},
	}

	err = cmd.Run(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := stdout.String()
	if !strings.Contains(output, "MY_VAR=hello") {
		t.Errorf("expected output to contain env var, got: %q", output)
	}
}

func TestRun_ExecutableNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	cmdDir := filepath.Join(tmpDir, "nonexistent")
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create dir: %v", err)
	}

	cmd, err := Load(folder.New(cmdDir))
	if err != nil {
		t.Fatalf("failed to load command: %v", err)
	}

	ctx := &RunContext{
		Args:   []string{},
		Stdin:  strings.NewReader(""),
		Stdout: &bytes.Buffer{},
		Stderr: &bytes.Buffer{},
	}

	err = cmd.Run(ctx)
	if err == nil {
		t.Error("expected error when executable not found")
	}
}
