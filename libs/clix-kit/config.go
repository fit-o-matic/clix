package kit

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds shared configuration for all clix plugins.
// Values are loaded from environment variables with the CLIX_ prefix.
type Config struct {
	// CLIName is the name of the parent CLI invoking the plugin.
	cliName string

	// CLIHome is the home directory of the parent CLI.
	cliHome string
}

// LoadConfig reads shared configuration from environment variables.
func LoadConfig() (*Config, error) {

	if os.Getenv("CLIX_CLI_NAME") == "" || os.Getenv("CLIX_CLI_HOME") == "" {
		return nil, fmt.Errorf("missing required CLIX_CLI_NAME or CLIX_CLI_HOME environment variables")
	}

	return &Config{
		cliName: os.Getenv("CLIX_CLI_NAME"),
		cliHome: os.Getenv("CLIX_CLI_HOME"),
	}, nil
}

func (c *Config) GetCliName() string {
	return c.cliName
}

func (c *Config) GetCliHome() string {
	return c.cliHome
}

func (c *Config) GetCliCacheDir() string {
	return filepath.Join(c.cliHome, "cache")
}
