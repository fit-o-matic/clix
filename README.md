# clix

CLI + "x" for extensible.

clix is a Go framework for building extensible CLI applications. Add functionality by dropping plugin executables into a plugins folder — each plugin is a standalone [12-factor app](https://12factor.net/) that communicates with clix through defined I/O interfaces.

## Project Structure

```
libs/
  clix/            # Core library
  clix-kit/        # CLI toolkit (help, version, cobra integration)
plugins/
  secret/          # Example plugin: secret management
```

This is a Go workspace monorepo. Each module has its own `go.mod` and can be built independently.

## Writing a Plugin

A plugin is a Go executable that uses the `clix-kit` library. Define your cobra commands and wire them into a `Plugin` struct:

```go
package main

import (
    kit "github.com/finkt/clix-kit"
    "github.com/spf13/cobra"
)

func main() {
    p := &kit.Plugin{
        Name:        "myplugin",
        Description: "does something useful",
        Version:     "0.1.0",
        Usage:       "myplugin [command]",
        Cmd:         newRootCmd(),
    }
    p.Execute()
}
```

The `Plugin` struct provides:

- **`-h` / `--help`** — Templated help output (customizable via `HelpTemplate`)
- **Error handling** — Errors written to stderr, non-zero exit on failure
- **Cobra integration** — Subcommand routing and argument validation via `Cmd`

## Build

```sh
go build ./libs/clix/...
go build ./libs/clix-kit/...
go build ./plugins/secret/...
```

## Test

```sh
go test ./...
go test -race ./...
```

## License

See [LICENSE](LICENSE).
