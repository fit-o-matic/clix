# clix Design Notes

- **Name**: clix = CLI + "x" for extensible
- **Extension model**: Extensible by adding plugins to a defined folder
- **Plugins**: Standalone executables that are 12-factor apps implementing special I/O interfaces
- **Plugin location**: `plugins/` directory in the repo; each plugin is its own Go module
- **Repo structure**: Go workspace monorepo; main library at `libs/clix/`
