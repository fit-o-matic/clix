package command

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/finkt/clix-kit/folder"
)

func TestNewRegistry(t *testing.T) {
	tmpDir := t.TempDir()
	parentFolder := folder.New(tmpDir)

	registry := NewRegistry(parentFolder)

	if registry == nil {
		t.Fatal("expected registry to be non-nil")
	}
	if registry.folder == nil {
		t.Fatal("expected registry folder to be non-nil")
	}
	if registry.cache == nil {
		t.Fatal("expected registry cache to be non-nil")
	}
}

func Test_loadCommands_EmptyRegistry(t *testing.T) {
	tmpDir := t.TempDir()
	registryDir := filepath.Join(tmpDir, "registry")
	if err := os.MkdirAll(registryDir, 0755); err != nil {
		t.Fatalf("failed to create registry dir: %v", err)
	}

	parentFolder := folder.New(tmpDir)
	registry := NewRegistry(parentFolder)

	commands, err := registry.loadCommands()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(commands) != 0 {
		t.Errorf("expected 0 commands, got %d", len(commands))
	}
}

func Test_loadCommands_WithCommands(t *testing.T) {
	tmpDir := t.TempDir()
	registryDir := filepath.Join(tmpDir, "registry")

	commandNames := []string{"build", "deploy", "test"}
	for _, name := range commandNames {
		cmdDir := filepath.Join(registryDir, name)
		if err := os.MkdirAll(cmdDir, 0755); err != nil {
			t.Fatalf("failed to create command dir %s: %v", name, err)
		}
	}

	parentFolder := folder.New(tmpDir)
	registry := NewRegistry(parentFolder)

	commands, err := registry.loadCommands()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(commands) != len(commandNames) {
		t.Errorf("expected %d commands, got %d", len(commandNames), len(commands))
	}
}

func Test_createManifest(t *testing.T) {
	tmpDir := t.TempDir()

	names := []string{"alpha", "beta"}
	var commands Commands
	for _, name := range names {
		cmdDir := filepath.Join(tmpDir, name)
		if err := os.MkdirAll(cmdDir, 0755); err != nil {
			t.Fatalf("failed to create dir: %v", err)
		}
		cmd, err := Load(folder.New(cmdDir))
		if err != nil {
			t.Fatalf("failed to load command: %v", err)
		}
		commands = append(commands, cmd)
	}

	registered := commands.createManifest()

	if len(registered.Summaries) != len(names) {
		t.Errorf("expected %d commands, got %d", len(names), len(registered.Summaries))
	}
	for i, name := range names {
		if registered.Summaries[i].Name != name {
			t.Errorf("expected command name %q, got %q", name, registered.Summaries[i].Name)
		}
	}
}

func TestGetManifest(t *testing.T) {
	tmpDir := t.TempDir()
	registryDir := filepath.Join(tmpDir, "registry")

	cmdDir := filepath.Join(registryDir, "mycommand")
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create command dir: %v", err)
	}

	parentFolder := folder.New(tmpDir)
	registry := NewRegistry(parentFolder)

	registered, err := registry.GetManifest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if registered == nil {
		t.Fatal("expected registered commands to be non-nil")
	}
}

func TestClearCache(t *testing.T) {
	tmpDir := t.TempDir()
	registryDir := filepath.Join(tmpDir, "registry")

	cmdDir := filepath.Join(registryDir, "mycommand")
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create command dir: %v", err)
	}

	parentFolder := folder.New(tmpDir)
	registry := NewRegistry(parentFolder)

	// Generate the manifest cache
	_, err := registry.GetManifest()
	if err != nil {
		t.Fatalf("failed to get manifest: %v", err)
	}

	// Clear the cache
	err = registry.ClearCache()
	if err != nil {
		t.Fatalf("failed to clear cache: %v", err)
	}
}

func TestGetCommand_Exists(t *testing.T) {
	tmpDir := t.TempDir()
	registryDir := filepath.Join(tmpDir, "registry")

	cmdDir := filepath.Join(registryDir, "mycommand")
	if err := os.MkdirAll(cmdDir, 0755); err != nil {
		t.Fatalf("failed to create command dir: %v", err)
	}

	parentFolder := folder.New(tmpDir)
	registry := NewRegistry(parentFolder)

	cmd, err := registry.GetCommand("mycommand")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmd == nil {
		t.Fatal("expected command to be non-nil")
	}
	if cmd.GetName() != "mycommand" {
		t.Errorf("expected command name %q, got %q", "mycommand", cmd.GetName())
	}
}

func TestGetCommand_NotExists(t *testing.T) {
	tmpDir := t.TempDir()
	registryDir := filepath.Join(tmpDir, "registry")
	if err := os.MkdirAll(registryDir, 0755); err != nil {
		t.Fatalf("failed to create registry dir: %v", err)
	}

	parentFolder := folder.New(tmpDir)
	registry := NewRegistry(parentFolder)

	cmd, err := registry.GetCommand("nonexistent")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmd != nil {
		t.Error("expected command to be nil for non-existent command")
	}
}
