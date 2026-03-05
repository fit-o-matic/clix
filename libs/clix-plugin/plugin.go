package plugin

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

// defaultHelpTemplate is the built-in help text template.
var defaultHelpTemplate = `{{.Name}} - {{.Description}}

Version: {{.Version}}

Usage:
  {{.Usage}}
`

// Plugin defines a clix plugin with standard lifecycle handling.
type Plugin struct {
	Name         string
	Version      string
	Description  string
	Usage        string
	HelpTemplate string
	Cmd          *cobra.Command
}

// helpTemplate returns the plugin's help template, falling back to the default.
func (p *Plugin) helpTemplate() string {
	if p.HelpTemplate != "" {
		return p.HelpTemplate
	}
	return defaultHelpTemplate
}

// Execute runs the plugin lifecycle: handles the "help" subcommand,
// delegates to Cmd, and exits non-zero on error.
func (p *Plugin) Execute() {
	if len(os.Args) > 1 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		t, err := template.New("help").Parse(p.helpTemplate())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: invalid help template: %v\n", err)
			os.Exit(1)
		}
		if err := t.Execute(os.Stdout, p); err != nil {
			fmt.Fprintf(os.Stderr, "error: rendering help: %v\n", err)
			os.Exit(1)
		}
		return
	}

	p.Cmd.SetArgs(os.Args[1:])
	if err := p.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
