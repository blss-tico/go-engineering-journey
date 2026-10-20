# Module 00 — Professional Go Development Environment

> **Go Engineering Journey** · Phase 0 — Professional Environment · Go 1.27

[Português (Brasil)](README.pt-BR.md) · [English curriculum](../../CURRICULUM.md) · [Repository home](../../README.md)

## Purpose

Build a reliable, reproducible Go development environment while learning how the Go toolchain, modules, packages, dependencies, compiled binaries, and Git work together. This module is a hands-on foundation—not a production application or an architecture exercise.

**Prerequisites:** A terminal, Git, and a supported Go 1.27 toolchain. Examples were exercised with Go 1.27.1 on Linux/amd64. Adapt commands and output paths to your operating system.

## Learning outcomes

After completing this module, you should be able to:

1. Inspect the installed toolchain and key Go environment variables.
2. Explain the differences among repositories, modules, packages, and source files.
3. Run, format, inspect, test, vet, and compile a small Go program.
4. Cross-compile a binary and inspect its target and build metadata.
5. Create and import a package within the same module.
6. Add an external module, understand `go.mod` and `go.sum`, and distinguish module graphs from package graphs.
7. Use Git to review, commit, synchronize, and document a reproducible lab.

## Repository layout

```text
modules/module-00/
├── README.md
├── README.pt-BR.md
└── hello-go/
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── greeting/
    │   └── greeting.go
    └── bin/               # Local build outputs; ignored by Git
```

`bin/` is generated locally and does not need to exist in the repository. The lab is intentionally one module with two packages.

## 1. Inspect your Go installation

```bash
go version
go env GOROOT GOPATH GOCACHE GOMODCACHE GOTOOLCHAIN
go env GOOS GOARCH CGO_ENABLED
```

| Setting | Meaning |
| --- | --- |
| `GOROOT` | Go distribution and standard library location |
| `GOPATH` | Workspace location, including the default module-cache parent |
| `GOCACHE` | Compilation cache |
| `GOMODCACHE` | Downloaded module cache |
| `GOTOOLCHAIN` | Toolchain selection behavior |
| `GOOS`, `GOARCH` | Target operating system and architecture |
| `CGO_ENABLED` | Whether cgo is enabled for the build |

**Reflection:** Why should the source repository not be placed inside `GOROOT`? How does `GOMODCACHE` differ from your project directory?

## 2. Understand the module boundary

From the repository root:

```bash
cd modules/module-00/hello-go
cat go.mod
go env GOMOD
go list -m
go list -f 'Name={{.Name}} ImportPath={{.ImportPath}}' .
```

The module path is:

```text
github.com/blss-tico/go-engineering-journey/modules/module-00/hello-go
```

A **repository** holds version-controlled files. A **module** is a versioned unit identified by a `go.mod` file. A **package** consists of Go files compiled together in a directory. A **source file** is one `.go` file. One module can contain many packages; an import path is a logical identity, not a machine-specific filesystem path.

## 3. Run, format, vet, and build

```bash
gofmt -w main.go greeting/greeting.go
go run .
go vet ./...
go test ./...
mkdir -p bin
go build -o bin/hello-linux-amd64 .
./bin/hello-linux-amd64
```

On a Linux/amd64 machine, the program should print a greeting, the target `linux/amd64`, and a title-cased Portuguese sentence. `go test ./...` may report `[no test files]`; that is not evidence of functional test coverage.

`go run .` builds and executes a temporary program; `go build -o ... .` writes an executable to a chosen path. `gofmt` is the canonical Go formatter. `go vet` performs selected static checks, not a proof of correctness.

## 4. Inspect binaries and cross-compile

```bash
file bin/hello-linux-amd64
go version -m bin/hello-linux-amd64
GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0 go build -o bin/hello-freebsd-amd64 .
file bin/hello-freebsd-amd64
go version -m bin/hello-freebsd-amd64
```

`GOOS` and `GOARCH` specify the **build target**. They are not necessarily the host OS and architecture. A FreeBSD executable built on Linux generally cannot be run directly on Linux. Cross-compilation with cgo may require an appropriate cross-compiler, so the example explicitly disables cgo.

`go version -m` can show the module path, toolchain/build settings, and Git revision information. When VCS metadata is available, `vcs.modified=true` indicates the build observed local changes; `false` indicates a clean state according to the toolchain's check. This metadata is useful for provenance but is not a complete supply-chain attestation.

## 5. Create a local package

**`hello-go/greeting/greeting.go`**

```go
package greeting

func Message() string {
	return "Go Professional Training"
}
```

**`hello-go/main.go`**

```go
package main

import (
	"fmt"
	"runtime"

	"github.com/blss-tico/go-engineering-journey/modules/module-00/hello-go/greeting"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Println(greeting.Message())
	fmt.Printf("target=%s/%s\n", runtime.GOOS, runtime.GOARCH)

	title := cases.Title(language.BrazilianPortuguese)
	fmt.Println(title.String("engenharia de software com go"))
}
```

The uppercase `Message` identifier is **exported**, making it accessible from another package. A lowercase `message` identifier would not be accessible from `main`.

The local `greeting` package does **not** need its own `go.mod`: its import path is the main module path plus `/greeting`. Adding a nested `go.mod` would create a separate module boundary.

Inspect package resolution:

```bash
go list ./...
go list -f 'Name={{.Name}} ImportPath={{.ImportPath}}' ./greeting
go list -f '{{.Imports}}' .
go list -f '{{.Module.Path}}' ./greeting
go list -f '{{.ImportPath}} {{.Standard}}' fmt
```

## 6. Add and maintain an external module

The sample imports two packages, `cases` and `language`, from **one external module**, `golang.org/x/text`.

```bash
go get golang.org/x/text@latest
go mod tidy
cat go.mod
cat go.sum
go list -m all
go mod verify
```

**Observed lab dependency:** `golang.org/x/text v0.42.0`. This is a record of the exercise, not an instruction to always install that version. `@latest` can resolve differently over time and may affect the module's minimum Go version. Review all changes to `go.mod` before committing.

- **`go.mod`** declares the main module, its minimum Go version, and module requirements. Go selects module versions using Minimal Version Selection (MVS).
- **`go.sum`** records checksums for module content and module metadata. It is **not** a traditional version lockfile.
- **`// indirect`** marks a module requirement that does not provide a package directly imported by the main module (with some additional module-graph cases).
- **`go get`** changes module requirements; **`go mod tidy`** reconciles requirements and checksums with the source and relevant tests.
- **`go mod verify`** checks integrity of downloaded modules in the cache; it does not audit their security or run your tests.

### Module graph vs. package graph

```bash
go mod graph
go list -deps -f '{{if not .Standard}}{{.ImportPath}}{{end}}' . | sed '/^$/d'
go list -f '{{.Imports}}' .
```

These commands answer different questions:

| Command | Shows |
| --- | --- |
| `go mod graph` | Module requirement edges |
| `go list -m all` | Selected modules in the build list |
| `go list -deps ...` | Packages needed to build the selected package |
| `go list -f '{{.Imports}}' .` | Direct imports of the main package |

In the recorded exercise, `x/text v0.42.0` required `x/tools v0.49.0`, `x/mod v0.41.0`, and `x/sync v0.23.0` in the **module graph**. Their packages did not appear in the filtered package-dependency output for this executable. A module in the build list is not necessarily compiled into the binary.

Directories named `internal` enforce import restrictions: code outside the permitted parent tree cannot import those packages directly.

## 7. Git workflow and build provenance

Keep local binaries out of commits by adding `bin/` to the repository `.gitignore`.

From the repository root:

```bash
git status
git diff
git add modules/module-00/hello-go/main.go \
  modules/module-00/hello-go/greeting/greeting.go \
  modules/module-00/hello-go/go.mod \
  modules/module-00/hello-go/go.sum
git diff --cached --check
git diff --cached
git commit -m "Add local package and external module dependency lab"
git push origin main
```

This is an illustrative workflow; **do not repeat a commit if the lab has already been committed**. `git diff` does not show untracked files until staged. `git status` distinguishes working-tree and index changes.

If a push is rejected because the remote advanced, inspect before integrating:

```bash
git fetch origin
git status
git log --oneline --graph --decorate -12 --all
git log --oneline HEAD..origin/main
git log --oneline origin/main..HEAD
```

For a clean working tree and unpublished local commits, `git rebase origin/main` may be appropriate. Resolve conflicts deliberately; never use `git push --force` as an automatic response to rejection.

## 8. Practice exercises

1. Explain `GOROOT`, `GOPATH`, `GOCACHE`, and `GOMODCACHE` without consulting notes.
2. Compare `go run .` and `go build -o bin/hello .`.
3. Build for Linux and FreeBSD and inspect both with `file` and `go version -m`.
4. Explain why `greeting` resolves locally without a separate `require` entry.
5. Temporarily change `Message` to `message`. Predict and explain the compiler error; then restore the code.
6. Compare `go mod graph`, `go list -m all`, and `go list -deps`.
7. Explain why `go.sum` is not a lockfile and why `// indirect` does not mean “not installed.”
8. Inspect Git history and verify that local and remote branches are synchronized.

## 9. Definition of Done

- [ ] Go 1.27 is installed and the environment variables are understood.
- [ ] `hello-go` has one module and at least two packages.
- [ ] `go run .` works and the output is understood.
- [ ] Formatting, vetting, and `go test ./...` complete without errors.
- [ ] Linux and FreeBSD build targets have been inspected.
- [ ] `go.mod`, `go.sum`, direct/indirect requirements, and checksums can be explained.
- [ ] The module and package graphs have been compared.
- [ ] `go mod verify` succeeds.
- [ ] Build outputs are ignored and source changes are committed.
- [ ] The repository is synchronized and `git status` is clean.
- [ ] The reflection below is completed in your own words.

## 10. What I learned

Complete this section after the lab:

- **Most important concept:** …
- **Most surprising observation:** …
- **A problem I encountered and how I solved it:** …
- **What I want to explore further:** …

## Next module

**Module 01 — Advanced Go Toolchain:** deeper exploration of compiler and linker behavior, toolchain commands, documentation, diagnostics, and reproducible development workflows.

---

**Official references:** [Go documentation](https://go.dev/doc/) · [Go modules reference](https://go.dev/ref/mod) · [Go command documentation](https://pkg.go.dev/cmd/go) · [Effective Go](https://go.dev/doc/effective_go)
