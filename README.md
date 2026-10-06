# Go Engineering Journey

> A hands-on journey from Go fundamentals to professional software engineering.

[Português (Brasil)](./README.pt-BR.md) · [Full Curriculum](./CURRICULUM.md)

## About

**Go Engineering Journey** is a long-term, hands-on learning project focused on becoming a professional Go software engineer.

The goal goes far beyond learning the syntax of the language. This journey combines Go with the computer science and engineering knowledge required to design, build, test, debug, secure, observe, optimize, deploy, and evolve production software.

The curriculum currently targets **Go 1.27** and is designed to evolve alongside the language and its ecosystem.

This repository is also a public learning record: exercises, labs, experiments, notes, benchmarks, mistakes, refactorings, and projects are all part of the journey.

## Core Principle

> **Learn idiomatic Go first; introduce abstraction only when the problem justifies it.**

Examples and exercises prioritize Go community conventions, simplicity, and the standard library.

Architectural approaches such as Clean Architecture, Hexagonal Architecture, Domain-Driven Design, and microservices are intentionally introduced later, after the foundations required to understand their benefits, costs, and trade-offs.

## What This Journey Covers

The curriculum progresses from fundamentals to advanced professional topics:

- Go language fundamentals and idiomatic Go
- Go toolchain, modules, packages, and dependency management
- Data Structures and Algorithms
- Memory, runtime, operating systems, and networking
- Testing, debugging, static analysis, profiling, and performance
- Git and professional collaboration
- SQL, PostgreSQL, NoSQL, Redis, and caching
- HTTP, REST, gRPC, GraphQL, WebSockets, and SSE
- Concurrency and parallelism
- Go runtime and scheduler
- RabbitMQ, NATS/JetStream, Kafka, and event-driven systems
- Application security and resilience engineering
- Observability with OpenTelemetry, Prometheus, and Grafana
- Distributed systems
- Docker and Kubernetes
- CI/CD, software supply chain, Platform Engineering, and SRE
- AI Engineering with Go
- Software Architecture and System Design
- Open Source, technical communication, and interview preparation

See the **[complete curriculum](./CURRICULUM.md)** for all phases, modules, projects, and professional checkpoints.

## Learning Roadmap

| Phase | Focus |
|---|---|
| 0 | Professional Environment |
| I | Go Fundamentals |
| II | Computer Science Foundations |
| III | Professional Go |
| IV | Data Engineering |
| V | Backend Engineering |
| VI | Concurrency & Parallelism |
| VII | Messaging & Event-Driven Systems |
| VIII | Production Engineering |
| IX | Distributed Systems |
| X | Cloud Native & Platform Engineering |
| XI | Go + AI Engineering |
| XII | Architecture, System Design & Career |

The complete journey contains **50 modules (0–49)** and a sequence of progressively more realistic projects.

## Projects

The practical work grows together with the curriculum instead of being a collection of disconnected tutorials.

| # | Project | Main Focus |
|---|---|---|
| 01 | Professional CLI | Go fundamentals, filesystem, errors, testing |
| 02 | DSA Library | Data structures, algorithms, generics |
| 03 | HTTP Server | `net/http`, networking, testing |
| 04 | REST API + PostgreSQL | Backend APIs and relational persistence |
| 05 | API + PostgreSQL + Redis | Caching and production patterns |
| 06 | gRPC Service | Protocol Buffers and RPC |
| 07 | Event-Driven Application | Messaging and delivery semantics |
| 08 | Concurrent Processing Engine | Goroutines, channels, synchronization |
| 09 | Distributed Services | Reliability and distributed systems |
| 10 | Cloud-Native Application | Containers, Kubernetes, observability |
| 11 | Go + AI Application | LLM integration, RAG/tools, production concerns |
| 12 | Capstone System | End-to-end production engineering |

Major portfolio projects may eventually move to dedicated repositories. This repository will remain the central learning record and index.

## Study Method

Topics generally follow this progression:

**Theory → Internals → Minimal Example → Implementation → Idiomatic Go Review → Common Mistakes → Exercise → Challenge → Professional Application → Review**

The goal is mastery, not speed. Topics may be expanded before moving to the next module.

## Current Progress

**Current phase:** Phase 0 — Professional Environment  
**Current module:** Module 0 — Professional Go 1.27 Environment

- [ ] Phase 0 — Professional Environment
- [ ] Phase I — Go Fundamentals
- [ ] Phase II — Computer Science Foundations
- [ ] Phase III — Professional Go
- [ ] Phase IV — Data Engineering
- [ ] Phase V — Backend Engineering
- [ ] Phase VI — Concurrency & Parallelism
- [ ] Phase VII — Messaging & Event-Driven Systems
- [ ] Phase VIII — Production Engineering
- [ ] Phase IX — Distributed Systems
- [ ] Phase X — Cloud Native & Platform Engineering
- [ ] Phase XI — Go + AI Engineering
- [ ] Phase XII — Architecture, System Design & Career

Progress will be updated as the journey advances.

## Repository Structure

The structure will evolve naturally as new needs appear. A possible direction is:

```text
go-engineering-journey/
├── README.md
├── README.pt-BR.md
├── CURRICULUM.md
├── CURRICULUM.pt-BR.md
├── docs/
├── notes/
├── exercises/
├── labs/
├── dsa/
├── projects/
└── ...
```

The repository will not be over-structured prematurely. Directories and abstractions should appear when there is a real reason for them.

## Engineering Philosophy

This repository intentionally preserves the evolution of the learning process.

Early code is expected to be simpler than later production-oriented code. It should not be retroactively overengineered merely to look sophisticated.

As new concepts are learned, previous solutions may be revisited to study:

- why a refactoring is useful;
- what problem an abstraction solves;
- what complexity it introduces;
- whether performance actually improves;
- whether the new design remains idiomatic Go.

Mistakes and failed experiments are useful engineering evidence when their lessons are documented.

## Go Version

The curriculum currently targets:

```text
Go 1.27
```

Version-specific behavior should be documented when relevant. The repository may be updated as new Go releases become appropriate for the training.

## Who Is This For?

This repository may be useful for:

- developers learning Go from the beginning;
- developers coming from another programming language;
- backend developers moving into Go;
- engineers interested in concurrency and distributed systems;
- DevOps, SRE, or Platform Engineers who want deeper Go knowledge;
- students preparing for professional Go roles;
- anyone who wants a structured path beyond language syntax.

## Contributing

This project is public not only to document my own learning, but also in the hope that it can help other students.

Contributions are welcome when they improve the learning material—for example:

- correcting technical mistakes;
- improving explanations;
- improving English or Portuguese translations;
- suggesting useful exercises;
- fixing non-idiomatic Go;
- improving tests or documentation;
- adding relevant references.

Please preserve the project's main principle:

> **Idiomatic Go first. Abstraction when justified.**

## Languages

- **English:** [README.md](./README.md) · [Curriculum](./CURRICULUM.md)
- **Português (Brasil):** [README.pt-BR.md](./README.pt-BR.md) · [Grade completa](./CURRICULUM.pt-BR.md)

## License

This repository is distributed under the terms of the [MIT License](./LICENSE).

## Final Goal

The objective is not simply to say:

> “I know Go.”

The objective is to be able to say:

> **“I can engineer production software with Go, understand the systems it runs on, and explain the trade-offs behind my decisions.”**

That is the journey.
