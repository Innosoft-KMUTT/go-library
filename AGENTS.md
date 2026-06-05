# AGENTS.md

Guidance for AI agents and developers working in this repository.

## What this is

`github.com/Innosoft-KMUTT/go-library` — a small, shared Go library for Innosoft-KMUTT projects (Go 1.22.3). The headline feature is the `datatype` package: nullable, JSON-friendly wrappers around `database/sql`'s `Null*` types. The rest of the repo is scaffolding: a `sample` package showing the expected method/test shape, and a `go-library-template` directory holding the Coder + Terraform workspace template used to spin up dev environments.

There is no application here. `main.go` exists only so the module compiles and is not meant to do anything.

## Layout

| Path                      | What it is                                                                 |
| ------------------------- | ------------------------------------------------------------------------- |
| `datatype/`               | Nullable SQL/JSON types (the actual library). One file + test per type.   |
| `sample/`                 | Example package (`SampleMethod`) showing the package + test convention.   |
| `main.go` / `main_test.go`| Trivial `package main` stub so the module builds; prints `"main"`.        |
| `go-library-template/`    | Coder workspace template (Terraform `.tf` + Docker build + nginx).        |
| `.devcontainer/`          | VS Code dev container (Dockerfile, zsh/p10k config, example config).      |
| `git_all.sh`              | Helper: `echo`es then `eval`s its args. Used by the Coder clone step.     |

## The `datatype` package

Each type embeds the matching `sql.Null*` struct and adds `MarshalJSON` / `UnmarshalJSON` so the value round-trips cleanly through both a SQL driver and JSON.

| Type          | Embeds              | JSON output when invalid | Notes                                              |
| ------------- | ------------------- | ------------------------ | -------------------------------------------------- |
| `NullString`  | `sql.NullString`    | `null`                   | Simplest: any JSON string → valid.                 |
| `NullBool`    | `sql.NullBool`      | `null`                   | Accepts bool, or string `"true"/"false"` (case-insensitive via `ParseBool`). |
| `NullInt32`   | `sql.NullInt32`     | `null`                   | Accepts number or numeric string.                  |
| `NullInt64`   | `sql.NullInt64`     | `null`                   | Accepts number or numeric string.                  |
| `NullFloat64` | `sql.NullFloat64`   | `null`                   | Accepts number or numeric string.                  |
| `NullUInt`    | `sql.NullInt64`     | `null`                   | Stored as `Int64`; rejects negative values with `"value must greater than zero"`. |

### Marshalling contract

- **Marshal:** if `Valid` is false → emit JSON `null`; otherwise emit the underlying value.
- **Unmarshal** (the numeric/bool types) tries the data two ways — as the native type (into a pointer, so JSON `null` becomes a `nil` pointer) and as a `string` — then decides:
  - JSON `null` or empty string `""` → `Valid = false`.
  - A native value present → set it, `Valid = true`.
  - Otherwise parse the string with the matching `strconv` parser (`ParseBool` / `ParseInt` / `ParseFloat`).
- This is why both `{"value": 5}` and `{"value": "5"}` are accepted — the types are deliberately lenient about string-vs-native input coming from JSON.

When extending or fixing these types, keep that two-attempt (native-pointer + string) unmarshal pattern consistent across all of them — recent history is a series of fixes making the `null` and empty-string cases line up across types.

## Conventions

- One type per file, named `null_<type>.go`, with a sibling `null_<type>_test.go`.
- Tests use `github.com/stretchr/testify/assert` and a small local `Test<Type>Type` struct with a single `Value` field to exercise JSON round-trips. Tests also `fmt.Println`/`Printf` their cases (verbose by design).
- `go.mod` lists testify (and its transitive deps) as `// indirect`; that's the only third-party dependency.
- Mixed Thai/English context lives in the team's Coder/devcontainer tooling; code and comments are English.

## Running it

```bash
go test ./...        # runs datatype, sample, and the main stub — all should pass
go build ./...       # builds the module
go vet ./...
```

There is no service to run. `go run .` just prints `main`.

## go-library-template (Coder workspace)

`go-library-template/` is **not** part of the Go module — it's the infrastructure template that provisions a Coder cloud workspace for developing this repo. It's only relevant when changing how dev environments are created.

- `main.tf` / `agent.tf` / `parameter.tf` / `workspace.tf` — Coder + `kreuzwerker/docker` Terraform. Builds an image from `build/Dockerfile` (based on `askforanywork/dev-coder`), runs code-server on port `13337`, and clones this repo into `~/go-library`.
- `parameter.tf` exposes workspace params: git branch (default `main`), clone command, frontend/backend run commands, and **Nginx External Port** (default `8081`).
- `nginx/default.conf` (and `default.dev.conf`) reverse-proxy `/api/ → 127.0.0.1:8080` and `/ → 127.0.0.1:3000`. These are generic frontend/backend ports for whatever app a workspace hosts — not used by this library itself.
- `build/init-on-build.sh` / `init-on-create.sh` set up oh-my-zsh, nvm, powerlevel10k, and fonts.

## `.devcontainer`

VS Code dev container alternative to Coder. `Dockerfile.dev` is based on `mcr.microsoft.com/devcontainers/go:1.22-bookworm`. Copy `devcontainer.example.json` to `devcontainer.json` (the real file is `.gitignore`d) to use it. It bind-mounts the workspace to `/go-library` and pulls in zsh/p10k/TabNine config from the host.

## Caveats

- `NullUInt` does not embed a `sql.NullUint*` (there isn't one) — it stores its value in `sql.NullInt64.Int64` and guards against negatives manually. Read/write it via the `Int64` field, not a `Uint` field.
- The numeric/bool `UnmarshalJSON` methods contain commented-out `fmt.Printf` debug lines from development — harmless, but don't mistake them for active logging.
- `main.go` / `sample` are placeholders. Don't build features into them; add new library code as its own package next to `datatype`.
- `git_all.sh` blindly `eval`s its arguments; it exists for the Coder startup script and is not a general-purpose utility.

<!-- last-updated: 2026-06-05T01:36:51+00:00 commit: 2f5353f -->
