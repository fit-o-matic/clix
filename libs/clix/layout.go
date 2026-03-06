package clix

import (
	"os"
	"path/filepath"
)

type Layout struct {
	cliName    string
	cliHomeDir string
}

func NewLayout(cliName string) (*Layout, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return &Layout{cliName: cliName, cliHomeDir: filepath.Join(home, "."+cliName)}, nil
}

func (l *Layout) GetCliName() string {
	return l.cliName
}

func (l *Layout) GetCliHomeDir() string {
	return l.cliHomeDir
}

func (l *Layout) GetCliCacheDir() string {
	return filepath.Join(l.cliHomeDir, ".cache", l.cliName)
}

func (l *Layout) GetCliConfigDir() string {
	return filepath.Join(l.cliHomeDir, ".config", l.cliName)
}

func (l *Layout) GetCliPluginsDir() string {
	return filepath.Join(l.cliHomeDir, ".plugins", l.cliName)
}
