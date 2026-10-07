# Module 00 — Professional Go 1.27 Environment

> Building a clean, reproducible foundation for the Go Engineering Journey.

[Português (Brasil)](./README.pt-BR.md) · [Main Curriculum](../../CURRICULUM.md)

## Status

**Phase:** 0 — Professional Environment  
**Module:** 00  
**Target version:** Go 1.27  
**Status:** In progress

---

## Learning Objectives

By the end of this module, you should be able to:

- validate a Go installation;
- explain at a high level how Go source code becomes an executable;
- use essential Go toolchain commands;
- inspect the Go environment;
- explain `GOROOT`, `GOPATH`, `GOCACHE`, and `GOMODCACHE`;
- understand the roles of `GOOS` and `GOARCH`;
- explain the difference between a module, package, and source file;
- initialize a Go module;
- build and run a simple Go program;
- explain `package main` and `func main`;
- format code using the official Go tooling;
- perform basic static checks with `go vet`;
- explore documentation using `go doc`;
- inspect packages with `go list`;
- understand the basic roles of `go.mod` and `go.sum`.

---

# 1. What Is Go?

Go is a statically typed, compiled programming language designed with a strong emphasis on simplicity, maintainability, fast builds, concurrency, and practical software engineering.

It was originally designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

Some characteristics that strongly influence how Go software is written include:

- simple language syntax;
- fast compilation;
- garbage collection;
- built-in concurrency primitives;
- composition instead of traditional class inheritance;
- interfaces implemented implicitly;
- a powerful standard library;
- standardized tooling;
- straightforward creation of native executables.

A recurring idea throughout this journey will be that learning Go is not only about learning its syntax. It is also about learning the conventions and engineering philosophy of its ecosystem.

---

# 2. Compiled Language

Consider:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

This is source code. The machine does not execute this text directly.

At a simplified level:

```text
Go source code
      │
      ▼
Go compiler/toolchain
      │
      ▼
Machine code / executable
      │
      ▼
Operating system executes it
```

This model will become more detailed later when we study compilers, linking, operating systems, memory, and the Go runtime.

For now, remember:

> Go normally produces native executables from source code.

---

# 3. Verify the Go Installation

Run:

```bash
go version
```

A typical result resembles:

```text
go version go1.27.x linux/amd64
```

The output tells us important information:

```text
go version go1.27.x linux/amd64
           └──────┘ └───┘ └───┘
            version    OS   arch
```

The exact patch version and platform may differ on your machine.

---

# 4. Discover the Go Toolchain

Run:

```bash
go help
```

Go provides a unified command-line tool for most everyday development tasks.

Important commands include:

```text
go run
go build
go install
go test
go fmt
go vet
go doc
go env
go list
go generate
go tool
go version
go mod
```

We will learn these commands gradually rather than memorizing them all at once.

The important idea is:

> The official Go toolchain is part of the language's development experience.

This standardization is one of Go's strengths.

---

# 5. Inspect the Environment

Run:

```bash
go env
```

This command displays Go's environment configuration.

Variables worth recognizing at this stage include:

```text
GOOS
GOARCH
GOROOT
GOPATH
GOMOD
GOCACHE
GOMODCACHE
GOTOOLCHAIN
```

You do not need to memorize every variable yet.

You should, however, begin understanding what these specific variables represent.

---

# 6. `GOOS` and `GOARCH`

`GOOS` identifies the target operating system.

Examples:

```text
linux
windows
darwin
freebsd
```

`GOARCH` identifies the target CPU architecture.

Examples:

```text
amd64
arm64
386
```

Inspect yours:

```bash
go env GOOS
go env GOARCH
```

These variables are particularly important because Go supports cross-compilation for many targets.

For example, from a compatible development environment:

```bash
GOOS=linux GOARCH=amd64 go build
```

or:

```bash
GOOS=windows GOARCH=amd64 go build
```

We will study cross-compilation properly later. At this point, the key idea is that the target operating system and architecture are explicit concepts in the Go toolchain.

---

# 7. `GOROOT`

Inspect it:

```bash
go env GOROOT
```

`GOROOT` points to the Go installation/toolchain tree.

Conceptually:

```text
GOROOT
  │
  ├── Go toolchain
  ├── standard library sources
  └── supporting Go installation files
```

In a normal modern installation, you generally should not manually configure `GOROOT`.

A useful mental model is:

> `GOROOT` is about the Go installation itself.

---

# 8. `GOPATH`

Inspect it:

```bash
go env GOPATH
```

Historically, `GOPATH` played a central role in organizing Go source code and dependencies.

Older Go projects commonly lived under a structure similar to:

```text
$GOPATH/
└── src/
    └── example.com/
        └── project/
```

Modern Go development uses **Go Modules**, so your projects no longer need to live inside `$GOPATH/src`.

`GOPATH` still exists and is still used by Go tooling for some purposes, but it should not be confused with the root directory of every modern Go project.

---

# 9. `GOCACHE`

Inspect it:

```bash
go env GOCACHE
```

`GOCACHE` identifies the build cache.

Go can reuse previously generated build artifacts when appropriate instead of rebuilding everything from scratch.

Conceptually:

```text
source code
    │
    ▼
compiler
    │
    ├── reusable build artifacts ──► GOCACHE
    │
    ▼
executable
```

This contributes to fast development cycles.

---

# 10. `GOMODCACHE`

Inspect it:

```bash
go env GOMODCACHE
```

`GOMODCACHE` is the module download cache.

It stores downloaded module content used as dependencies.

Do not confuse it with `GOCACHE`:

| Variable | Main purpose |
|---|---|
| `GOCACHE` | Cache generated build artifacts |
| `GOMODCACHE` | Cache downloaded Go modules |

This distinction will become clearer when we begin using external modules.

---

# 11. Go Modules

A **module** is a collection of related Go packages versioned together.

Modern Go projects generally use a `go.mod` file to define their module.

Create a directory:

```bash
mkdir hello-go
cd hello-go
```

Initialize a module:

```bash
go mod init example.com/hello-go
```

This creates:

```text
hello-go/
└── go.mod
```

A minimal `go.mod` resembles:

```go
module example.com/hello-go

go 1.27
```

The module path is an identifier for the module.

For a real project, it commonly corresponds to the repository location, for example:

```text
github.com/username/project
```

We will study module paths, semantic versioning, dependency resolution, and module commands in much greater depth later.

---

# 12. Module vs. Package vs. Source File

These concepts must not be confused.

A simplified hierarchy is:

```text
Module
│
├── Package A
│   ├── file1.go
│   └── file2.go
│
└── Package B
    ├── file3.go
    └── file4.go
```

A **source file** is a `.go` file.

A **package** groups Go source files that belong together.

A **module** contains one or more packages that are versioned together.

For our first program:

```text
hello-go/                    ← module directory
├── go.mod                   ← module definition
└── main.go                  ← Go source file
                              └─ belongs to package main
```

Later modules will refine this mental model.

---

# 13. Your First Program

Create:

```text
main.go
```

with:

```go
package main

import "fmt"

func main() {
	fmt.Println("Go Professional Training")
	fmt.Println("Module 00")
	fmt.Println("Environment ready!")
}
```

The project now looks like:

```text
hello-go/
├── go.mod
└── main.go
```

Notice how small the structure is.

We are deliberately **not** creating directories such as:

```text
cmd/
internal/
pkg/
services/
repositories/
controllers/
```

There is currently no problem that requires them.

This is intentional.

> Structure should grow from real requirements, not from copied templates.

---

# 14. `package main`

Every Go source file begins with a package declaration.

Our file contains:

```go
package main
```

The package named `main` has a special role: it is used to define executable programs.

An executable `main` package needs:

```go
func main()
```

This function is the entry point of the program.

Conceptually:

```text
Operating system starts executable
             │
             ▼
        package main
             │
             ▼
          main()
             │
             ▼
       program logic
```

---

# 15. Imports

Our program contains:

```go
import "fmt"
```

`fmt` is a package from the Go standard library.

We use:

```go
fmt.Println(...)
```

to write formatted output.

The capital `P` in `Println` is meaningful.

In Go, identifiers beginning with an uppercase letter are exported from their package.

Go does not use keywords such as:

```text
public
private
protected
```

for this purpose.

We will study visibility and package APIs in depth later.

---

# 16. Run the Program

From the module directory:

```bash
go run .
```

Expected output:

```text
Go Professional Training
Module 00
Environment ready!
```

The dot means the package in the current directory.

At a high level, `go run` builds what is necessary and runs the resulting program for you.

It is convenient during development.

---

# 17. Build the Program

Now run:

```bash
go build
```

Go creates an executable for the `main` package.

On Linux/macOS, depending on the module/directory name, you may then run something similar to:

```bash
./hello-go
```

On Windows, the generated executable will normally use `.exe`.

Expected output:

```text
Go Professional Training
Module 00
Environment ready!
```

The conceptual difference is:

```text
go run .
   │
   ├── build temporary executable
   └── execute it

go build
   │
   └── create executable artifact
```

This distinction becomes important in professional build and deployment workflows.

---

# 18. Formatting Go Code

Go has an official formatting standard.

Run:

```bash
gofmt -w main.go
```

or:

```bash
go fmt ./...
```

The first command directly formats the specified file.

The second formats packages selected by the pattern `./...`.

One of Go's cultural strengths is that formatting is largely standardized.

Instead of teams spending significant time debating formatting style, Go provides a canonical formatter.

> Formatting is tooling, not personal taste.

---

# 19. Basic Static Checks with `go vet`

Run:

```bash
go vet ./...
```

`go vet` analyzes Go code for suspicious constructs that may indicate mistakes.

It is useful to distinguish several tools conceptually:

```text
gofmt / go fmt
    │
    └── formatting

go vet
    │
    └── suspicious constructs / static checks

go test
    │
    └── execute tests

Staticcheck and other linters
    │
    └── additional static analysis
```

These tools overlap somewhat, but they do not serve exactly the same purpose.

We will study them individually later.

---

# 20. Documentation from the Terminal

Go documentation can be explored without leaving the terminal.

Try:

```bash
go doc fmt
```

Then:

```bash
go doc fmt.Println
```

This is a valuable habit.

Instead of immediately searching the internet for every API, learn to inspect the documentation available through the toolchain.

---

# 21. Inspect Packages with `go list`

Try:

```bash
go list std
```

This lists packages in the standard library.

You will encounter packages such as:

```text
fmt
errors
io
os
strings
bytes
net
net/http
context
sync
time
testing
```

You are not expected to learn them now.

The important lesson is that Go ships with a substantial standard library and provides tooling to inspect it.

---

# 22. `go.mod` and `go.sum`

At this stage, your project may only need:

```text
go.mod
```

As dependencies are introduced, you will commonly also see:

```text
go.sum
```

A simplified distinction is:

```text
go.mod
  └── declares the module and dependency requirements

go.sum
  └── records cryptographic checksums used to verify module content
```

`go.sum` should not be described simply as Go's equivalent of a lock file. Its role is different.

We will study dependency management properly in Module 1 and later professional Go modules.

---

# 23. First Laboratory

Create the following inside this repository:

```text
modules/
└── module-00/
    ├── README.md
    ├── README.pt-BR.md
    └── hello-go/
        ├── go.mod
        └── main.go
```

Inside `hello-go`, initialize the module.

A repository-aware module path could be:

```bash
go mod init github.com/blss-tico/go-engineering-journey/modules/module-00/hello-go
```

Then create:

```go
package main

import "fmt"

func main() {
	fmt.Println("Go Professional Training")
	fmt.Println("Module 00")
	fmt.Println("Environment ready!")
}
```

Run the complete validation sequence:

```bash
go fmt ./...
go vet ./...
go run .
go build
```

Then execute the generated binary.

### Expected program output

```text
Go Professional Training
Module 00
Environment ready!
```

---

# 24. Environment Investigation Exercise

Use the Go toolchain to discover the values on **your machine**:

```bash
go version
go env GOOS
go env GOARCH
go env GOROOT
go env GOPATH
go env GOCACHE
go env GOMODCACHE
go env GOTOOLCHAIN
```

Do not copy values from this document.

The purpose of the exercise is to investigate your actual development environment.

### My environment

Fill this in after running the commands:

```text
Go version:
GOOS:
GOARCH:
GOROOT:
GOPATH:
GOCACHE:
GOMODCACHE:
GOTOOLCHAIN:
```

Avoid committing machine-specific secrets or sensitive paths if your environment contains any.

---

# 25. Exercises

## Exercise 1 — Explain the hierarchy

Without looking at the previous section, explain:

```text
Module → Package → Source file
```

Write the explanation in your own words.

## Exercise 2 — Explain the environment

Explain the difference between:

```text
GOROOT
GOPATH
GOCACHE
GOMODCACHE
```

## Exercise 3 — `run` vs. `build`

Explain what practical difference you observe between:

```bash
go run .
```

and:

```bash
go build
```

## Exercise 4 — Explore the standard library

Use:

```bash
go doc
```

to investigate at least three standard-library packages.

Suggested starting points:

```text
fmt
os
strings
```

Write one sentence explaining the primary purpose of each package.

## Exercise 5 — Explore the toolchain

Run:

```bash
go help
```

Choose three commands we have not studied deeply yet and write a short hypothesis about what each one does.

Do not worry about being perfect. We will revisit them.

---

# 26. Challenge — Cross-Compilation Observation

Do not treat this as a production deployment exercise yet.

First inspect:

```bash
go env GOOS GOARCH
```

Then research the available target combinations through the Go toolchain:

```bash
go tool dist list
```

Observe the format:

```text
GOOS/GOARCH
```

Choose one target different from your current environment and identify what would change in the build command.

The purpose is to understand the concept, not to master cross-compilation yet.

---

# 27. Checkpoint

Before considering Module 00 complete, you should be able to answer these questions without simply repeating definitions:

1. What does it mean for Go to be compiled?
2. What information does `go version` provide?
3. What are `GOOS` and `GOARCH`?
4. What is `GOROOT`?
5. Why is `GOPATH` less central to project layout than it was historically?
6. What is the difference between `GOCACHE` and `GOMODCACHE`?
7. What is a Go module?
8. What is a Go package?
9. What is a Go source file?
10. What does `go mod init` do?
11. Why does an executable program use `package main`?
12. What is the role of `func main()`?
13. What is the practical difference between `go run .` and `go build`?
14. Why is `gofmt` important to Go culture?
15. What does `go vet` attempt to detect?
16. How can you inspect documentation from the terminal?
17. What is the basic purpose of `go.mod`?
18. What is the basic purpose of `go.sum`?
19. Why are we not creating a complex directory structure yet?

If any answer is unclear, that topic deserves review before moving on.

---

# 28. Engineering Notes

## Why are we starting with a tiny project?

Because architecture should solve actual problems.

Our program currently has one responsibility and only a few lines of code. Adding several layers and directories would increase cognitive overhead without solving a real engineering problem.

Later, when our applications develop genuine boundaries and complexity, we will have concrete reasons to introduce structure.

This allows us to observe an important progression:

```text
simple problem
     │
     ▼
simple solution
     │
     ▼
new requirements
     │
     ▼
new constraints
     │
     ▼
justified structure
```

That progression is more educational than beginning with a large template whose abstractions we do not yet understand.

## Standard library first

Throughout the journey, we will frequently investigate whether the Go standard library already provides what we need before adding external dependencies.

This does **not** mean third-party libraries are bad.

It means dependencies should be selected deliberately.

---

# 29. What I Learned

Complete this section after the laboratory.

### Concepts that became clear

- 
- 
- 

### Concepts I still need to review

- 
- 
- 

### Something that surprised me

- 

### Questions for the next lesson

- 

---

# 30. Commands Used in This Module

```bash
go version
go help
go env
go env GOOS
go env GOARCH
go env GOROOT
go env GOPATH
go env GOCACHE
go env GOMODCACHE
go env GOTOOLCHAIN
go mod init <module-path>
go run .
go build
gofmt -w main.go
go fmt ./...
go vet ./...
go doc fmt
go doc fmt.Println
go list std
go tool dist list
```

Do not memorize this list mechanically. Learn what problem each command solves.

---

# 31. Definition of Done

Module 00 is complete when:

- [ ] Go 1.27 is installed and validated.
- [ ] The environment investigation is complete.
- [ ] `hello-go` has its own `go.mod`.
- [ ] The program runs successfully with `go run .`.
- [ ] The program builds successfully with `go build`.
- [ ] The generated executable runs successfully.
- [ ] `go fmt ./...` completes successfully.
- [ ] `go vet ./...` completes successfully.
- [ ] The exercises have been answered.
- [ ] The checkpoint questions can be explained in your own words.
- [ ] The **What I Learned** section has been completed.

---

# Next

After completing and reviewing this module, continue to:

**Module 01 — Go Toolchain**

There we will move from simply using a few commands to understanding the Go development toolchain more systematically.
