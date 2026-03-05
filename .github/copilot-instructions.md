# Project Guidelines

## Overview

clix is a Go monorepo providing extensible building blocks for CLI applications. This project follows open source best practices and is intended for public contribution.

The repo uses a Go workspace (`go.work`) to manage multiple modules.

## Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go) and the [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments).
- Use `gofmt` / `goimports` for formatting — no exceptions.
- Prefer stdlib over third-party packages unless there is a clear, justified need.
- Keep exported API surfaces small; unexport anything that doesn't need to be public.
- Use meaningful, concise names following Go conventions: short variable names in small scopes, descriptive names for exported identifiers.
- Avoid `init()` functions. Prefer explicit initialization.

## Error Handling

- Always handle errors explicitly — never discard with `_` unless there is a documented reason.
- Use `fmt.Errorf` with `%w` for wrapping errors to preserve the error chain.
- Return errors to the caller; avoid `log.Fatal` or `os.Exit` outside of `main`.
- Prefer sentinel errors or custom error types over string matching.

## Project Structure

This is a Go workspace monorepo. Each module has its own `go.mod`.

- `libs/clix/` — The clix library (`github.com/finkt/clix`).
- `cmd/` — Application entry points (future modules).
- `internal/` — Private packages not importable by external projects.
- Within each module, follow [Standard Go Project Layout](https://github.com/golang-standards/project-layout) conventions.
- Keep `main.go` minimal — parse flags/config, wire dependencies, call `run()`.
- When adding a new module, register it with `go work use ./path/to/module`.

## Testing

- Write table-driven tests using `t.Run` subtests.
- Use the `testing` package from stdlib; avoid test framework dependencies.
- Test files live alongside the code they test (`foo_test.go` next to `foo.go`).
- Use `testdata/` directories for fixtures.
- Aim for meaningful coverage of behavior, not line-count metrics.
- Run tests: `go test ./...`
- Run tests with race detector: `go test -race ./...`

## Dependencies

- Use Go modules (`go.mod` / `go.sum`) per module.
- Pin dependencies to specific versions; review upgrades deliberately.
- Run `go mod tidy` in each affected module before committing.
- Minimize external dependencies — every dependency is a long-term maintenance cost.
- The `go.work` file is committed; each module must still build independently via its own `go.mod`.

## Documentation

- Every exported type, function, and package must have a doc comment.
- Package comments go in `doc.go` for non-trivial packages.
- Write doc comments as complete sentences starting with the identifier name.
- Keep README.md current with build, install, and usage instructions.

## CLI Conventions

- Use consistent `--flag-name` kebab-case for flags.
- Always provide `--help` output that is clear and complete.
- Write to stdout for normal output, stderr for errors and diagnostics.
- Exit with code 0 on success, non-zero on failure.
- Support `--version` to print the build version.

## Open Source Readiness

- All files must include the project license header where required.
- Commit messages follow [Conventional Commits](https://www.conventionalcommits.org/) format.
- Public APIs must be stable or clearly marked as experimental.
- Avoid logging secrets, tokens, or sensitive data.
- Keep the CI pipeline green — never merge with failing tests or lint errors.

## Build & CI

- Build: `go build ./...`
- Lint: `golangci-lint run`
- Vet: `go vet ./...`
