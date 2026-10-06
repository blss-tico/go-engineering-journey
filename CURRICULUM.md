# Go Engineering Journey — Curriculum

> From Go fundamentals to professional software engineering.

**Language:** English | [Português (Brasil)](./CURRICULUM.pt-BR.md)

## About this curriculum

This repository documents a long-term, hands-on journey to become a professional Go software engineer. The goal is not only to learn Go syntax, but to develop the computer science, software engineering, backend, infrastructure, distributed systems, security, observability, cloud-native, and AI engineering skills expected in real-world work.

The curriculum targets **Go 1.27** and should evolve as the language and ecosystem evolve.

## Guiding principles

1. **Idiomatic Go first.** Examples, exercises, and early projects follow Go community conventions and favor simple, idiomatic solutions.
2. **Standard library first.** Learn what Go already provides before introducing frameworks or third-party abstractions.
3. **Understand before abstracting.** Clean Architecture, Hexagonal Architecture, DDD, repositories, services, and similar approaches are introduced only after the foundations needed to evaluate their trade-offs.
4. **Theory + implementation.** Important concepts are studied conceptually and then implemented or investigated in Go.
5. **Production mindset.** Testing, debugging, security, performance, observability, failure handling, and operations are part of engineering.
6. **Evidence-driven optimization.** Measure before optimizing.
7. **Build in public.** Exercises, labs, documentation, Git history, and projects form a visible learning record and portfolio.
8. **AI as a tool, not a substitute for understanding.** AI-assisted development is used critically, with human review and verification.

## Learning path

| Phase | Focus | Outcome |
|---|---|---|
| 0 | Professional Environment | Productive Go development environment |
| I | Go Fundamentals | Strong command of the language |
| II | Computer Science | DSA, memory, OS, and networking foundations |
| III | Professional Go | Testing, debugging, quality, performance, Git |
| IV | Data Engineering | SQL, PostgreSQL, NoSQL, Redis |
| V | Backend Engineering | HTTP, REST, gRPC, GraphQL, API engineering |
| VI | Concurrency & Parallelism | Go concurrency model and runtime |
| VII | Messaging & Event-Driven Systems | Brokers, delivery semantics, event patterns |
| VIII | Production Engineering | Security, resilience, observability, performance |
| IX | Distributed Systems | Consistency, reliability, failure handling |
| X | Cloud Native & Platform Engineering | Containers, Kubernetes, CI/CD |
| XI | AI Engineering | Building AI-enabled systems with Go |
| XII | Architecture, System Design & Career | Advanced design and professional readiness |

---

## Phase 0 — Professional Environment

### Module 0 — Professional Go 1.27 Environment
Installing and validating Go; Linux development environment; `GOROOT`, `GOPATH`, `GOCACHE`, `GOMODCACHE`; `GOOS` and `GOARCH`; modules; editor/IDE; terminal workflow; repository setup; first build and execution.

### Module 1 — Go Toolchain
`go run`, `go build`, `go install`, `go test`, `go fmt`/`gofmt`, `go vet`, `go doc`, `go env`, `go list`, `go generate`, `go tool`, and module/dependency commands.

## Phase I — Go Fundamentals

### Module 2 — Language Fundamentals
Packages; variables and constants; built-in types; operators; control flow; zero values; scope.

### Module 3 — Go Type System
Static typing; defined and alias types; conversions; type inference; comparability; type identity.

### Module 4 — Arrays, Slices and Maps
Arrays; slice internals; length/capacity; `append`; `copy`; maps; `make`; common pitfalls.

### Module 5 — Strings, Bytes, Runes and Unicode
Strings; UTF-8; `byte`; `rune`; Unicode; `strings`; `bytes`.

### Module 6 — Structs, Pointers and Memory Basics
Structs; pointers; value semantics; `new`; `make`; basic stack/heap mental model.

### Module 7 — Functions, Methods, Interfaces and Composition
Functions; multiple returns; variadic functions; closures; `defer`; methods/receivers; method sets; interfaces; implicit implementation; embedding and composition.

### Module 8 — Errors and Generics
Idiomatic errors; wrapping; `errors.Is`; `errors.As`; custom errors; `panic`/`recover`; type parameters; constraints; inference; generic data structures; relevant Go 1.27 features.

## Phase II — Computer Science Foundations

### Module 9 — Data Structures
Dynamic arrays; linked lists; stacks; queues/deques; hash tables; sets; heaps/priority queues; trees; BSTs; balanced-tree concepts; tries; graphs.

### Module 10 — Algorithms
Time/space complexity; Big-O; searching; sorting; recursion; divide and conquer; greedy; backtracking; dynamic programming; BFS/DFS; shortest paths; topological sorting; interview-style problem solving.

### Module 11 — Memory and Runtime Foundations
Stack vs. heap; escape analysis; allocations; garbage collection; memory layout; alignment; locality; allocation pressure.

### Module 12 — Operating Systems for Go Engineers
Processes/threads; scheduling; context switching; virtual memory; syscalls; file descriptors; signals; filesystems; blocking/non-blocking I/O; `epoll` concepts; namespaces; cgroups.

### Module 13 — Networking Foundations
TCP/IP; IP; TCP/UDP; DNS; sockets; TLS; HTTP/1.1; HTTP/2; HTTP/3 concepts; keep-alive; connection pooling; timeouts.

## Phase III — Professional Go

### Module 14 — Packages, Modules and API Design
Package cohesion; public/internal APIs; `internal`; module boundaries; dependencies; compatibility; versioning; package cycles; idiomatic package design.

### Module 15 — Testing
`testing`; table-driven tests; subtests; helpers; test doubles; integration/E2E/contract tests; fuzzing; property-based concepts; coverage; race-enabled testing.

### Module 16 — Debugging
Delve; breakpoints; conditional breakpoints; stack traces; goroutine inspection; concurrency debugging; race detector; remote debugging concepts.

### Module 17 — Code Quality and Static Analysis
`gofmt`; `go vet`; Staticcheck; linters; code review; dependency hygiene; Go code smells; maintaining idiomatic Go.

### Module 18 — Benchmarking, Profiling and Performance
Benchmarks; CPU/heap/allocation profiles; `pprof`; blocking/mutex profiles; goroutine leak analysis; flame graph concepts; PGO; performance methodology.

### Module 19 — Professional Git and Collaboration
Git fundamentals/internals; branches; merge/rebase; conflicts; commits; pull requests; reviews; tags; semantic versioning; releases; collaborative workflows.

## Phase IV — Data Engineering

### Module 20 — Database Theory
Relational model; modeling; normalization/denormalization; keys/constraints; ACID; transactions; isolation; MVCC; locks; deadlocks; indexes; B-trees; query planning.

### Module 21 — PostgreSQL with Go
SQL; `database/sql`; `pgx`; pools; prepared statements; transactions; migrations; `EXPLAIN`; query optimization; errors; database testing.

### Module 22 — NoSQL and Data-Store Selection
Key-value; document; wide-column and time-series concepts; MongoDB case study; SQL vs. NoSQL trade-offs; storage selection.

### Module 23 — Redis and Caching
Redis; cache-aside; TTL; eviction; invalidation; distributed caches; stampede; penetration; serialization; when not to cache.

## Phase V — Backend Engineering

### Module 24 — HTTP in Depth
HTTP semantics; methods; status codes; headers; cookies; content negotiation; compression; connections; timeouts; `net/http`; servers and clients.

### Module 25 — REST APIs
Resource modeling; routing; middleware; validation; errors; pagination; filtering; sorting; versioning; OpenAPI; documentation; testing.

### Module 26 — Advanced API Engineering
Idempotency; rate limiting; backward compatibility; API gateways; webhooks; WebSockets; SSE; streaming; API lifecycle.

### Module 27 — gRPC and Protocol Buffers
Protobuf; services; unary/server/client/bidirectional streaming; interceptors; deadlines; cancellation; errors; API evolution.

### Module 28 — GraphQL
Schemas; queries; mutations; resolvers; DataLoader; N+1; authorization; subscriptions; trade-offs.

**Protocol comparison:** REST vs. gRPC vs. GraphQL vs. WebSockets vs. SSE.

## Phase VI — Concurrency and Parallelism

### Module 29 — Go Concurrency
Goroutines; channels; buffering; `select`; mutex/RWMutex; atomics; WaitGroup; `context`; cancellation; races; deadlocks.

### Module 30 — Concurrency Patterns
Worker pools; pipelines; fan-out/fan-in; producer/consumer; semaphores; rate limiting; bounded concurrency; graceful shutdown; backpressure.

### Module 31 — Go Runtime and Scheduler
Goroutine internals; G/M/P; scheduler; work stealing; preemption; network poller; blocking syscalls; GC interaction.

### Module 32 — Parallelism and Low-Level Performance
Concurrency vs. parallelism; CPU-bound vs. I/O-bound; `GOMAXPROCS`; multicore; contention; atomics; false sharing; CPU caches; scalability measurement.

## Phase VII — Messaging and Event-Driven Systems

### Module 33 — Messaging Theory
Queues; topics; producers/consumers; pub/sub; acknowledgements; ordering; delivery semantics; retries; DLQs; idempotency; backpressure; schema evolution.

### Module 34 — RabbitMQ, NATS/JetStream and Kafka
Architecture; Go clients; delivery/persistence; consumer groups; streams; operations; technology selection; Transactional Outbox; Saga; CQRS; Event Sourcing concepts.

## Phase VIII — Production Engineering

### Module 35 — Application Security
Threat modeling; OWASP; TLS/mTLS; authentication/authorization; password hashing; JWT; OAuth 2.0/OIDC; secrets; validation; SQL injection; SSRF; rate limiting; dependency/supply-chain security; secure defaults.

### Module 36 — Resilience Engineering
Timeouts; deadlines; retries; exponential backoff; jitter; circuit breakers; bulkheads; load shedding; backpressure; graceful degradation/shutdown.

### Module 37 — Observability
Structured logs; metrics; distributed tracing; OpenTelemetry; Prometheus; Grafana; RED; USE; SLIs/SLOs; alerting; context propagation.

### Module 38 — Production Performance Engineering
Load/stress/soak testing; throughput; latency; percentiles; capacity planning; production profiling; bottleneck analysis; regressions.

## Phase IX — Distributed Systems

### Module 39 — Distributed Systems Foundations
Failure models; partitions; CAP; consistency models; replication; partitioning; logical time concepts; consensus; leader election; eventual consistency; distributed transactions.

### Module 40 — Distributed Reliability Patterns
Idempotency; deduplication; distributed locks; Saga; Transactional Outbox; retry safety; circuit breakers; failure recovery; partial failure; duplicate/delayed messages.

## Phase X — Cloud Native and Platform Engineering

### Module 41 — Containers and Docker
Containers; images; Dockerfiles; multi-stage/minimal images; networking; volumes; resource limits; signals; security; graceful shutdown; health checks.

### Module 42 — Kubernetes
Pods; Deployments; Services; ConfigMaps; Secrets; probes; requests/limits; HPA; Ingress; Helm; controllers/operators; Go in Kubernetes.

### Module 43 — CI/CD and Software Supply Chain
CI/CD; GitHub Actions; tests/static analysis; builds; registries; SBOM; vulnerability scanning; signing concepts; releases; deployment strategies; Platform Engineering and SRE concepts.

## Phase XI — Go + AI Engineering

### Module 44 — AI Engineering for Go Developers
LLM fundamentals; tokens/context; model APIs; streaming; structured outputs; embeddings; vector search; RAG; tool/function calling; agents; evaluation; guardrails; prompt-injection awareness; observability; rate limits; cost; Go AI services; AI-assisted engineering; reviewing generated Go.

## Phase XII — Architecture, System Design and Career

### Module 45 — Software Architecture
After strong Go foundations: modular monoliths; microservices; Clean Architecture; Hexagonal/Ports and Adapters; DDD; Vertical Slice concepts; event-driven architecture; dependency boundaries; trade-offs; overengineering; adapting patterns to idiomatic Go.

### Module 46 — System Design
Functional/non-functional requirements; capacity estimation; throughput; latency; bandwidth; storage; load balancing; caching; partitioning; replication; fault tolerance; availability; scalability; trade-offs; design interviews.

### Module 47 — Open Source and Reading Real Go
Large codebases; standard-library source; established Go projects; issues; PRs; contribution workflow; RFC/design documents; open-source contribution.

### Module 48 — Technical Documentation and Communication
READMEs; package/API docs; ADRs; RFCs; architecture diagrams; technical writing; technical English; communicating trade-offs.

### Module 49 — Career and Interviews
Go résumé; GitHub portfolio; LinkedIn; Go/DSA/SQL/concurrency/debugging/System Design interviews; live coding; behavioral communication; skill-gap analysis.

---

# Projects

| # | Project | Main focus |
|---|---|---|
| 01 | Professional CLI | Fundamentals, filesystem, errors, testing |
| 02 | DSA Library | Data structures, algorithms, generics |
| 03 | HTTP Server | `net/http`, networking, testing |
| 04 | REST API + PostgreSQL | API and relational persistence |
| 05 | API + PostgreSQL + Redis | Caching and production patterns |
| 06 | gRPC Service | Protobuf and RPC |
| 07 | Event-Driven Application | Messaging and delivery semantics |
| 08 | Concurrent Processing Engine | Goroutines, channels, synchronization |
| 09 | Distributed Services | Reliability and distributed systems |
| 10 | Cloud-Native Application | Containers, Kubernetes, observability |
| 11 | Go + AI Application | LLM integration, RAG/tools, production concerns |
| 12 | Capstone System | End-to-end production engineering |

Major portfolio projects may move to their own repositories while this repository keeps the learning record, labs, notes, and links.

# Professional Checkpoints

These checkpoints measure technical competency, not job title or years of experience.

### Checkpoint 1 — Go Fundamentals
Target: solid entry-level Go foundations.

### Checkpoint 2 — Go + HTTP + SQL + Git
Target: junior backend readiness.

### Checkpoint 3 — Concurrency + Databases + APIs + Testing
Target: strong junior / early mid-level technical breadth.

### Checkpoint 4 — Messaging + Observability + Docker + Performance
Target: mid-level production engineering skills.

### Checkpoint 5 — Distributed Systems + Kubernetes + System Design
Target: advanced/senior-level technical competencies.

### Checkpoint 6 — Runtime + Performance + Architecture + Platform + Open Source
Target: advanced Go engineering depth.

---

# Study Method

Each topic should normally progress through:

1. **Theory** — understand the concept and vocabulary.
2. **Internals** — understand what Go or the underlying system is doing.
3. **Minimal example** — isolate the idea with the smallest useful program.
4. **Hands-on implementation** — write the code.
5. **Idiomatic Go review** — compare the solution with Go community conventions.
6. **Common mistakes** — learn failure modes and misleading approaches.
7. **Exercise** — reinforce the concept.
8. **Challenge** — solve a less guided problem.
9. **Professional application** — connect the concept to real software.
10. **Review** — explain the concept without relying on the example.

The goal is not to rush through modules. A topic may be expanded before moving forward.

---

# Progress Tracking

Use the following markers:

- `[ ]` Not started
- `[~]` In progress
- `[x]` Completed
- `[R]` Review needed

## Phase progress

- [ ] Phase 0 — Professional Environment
- [ ] Phase I — Go Fundamentals
- [ ] Phase II — Computer Science Foundations
- [ ] Phase III — Professional Go
- [ ] Phase IV — Data Engineering
- [ ] Phase V — Backend Engineering
- [ ] Phase VI — Concurrency and Parallelism
- [ ] Phase VII — Messaging and Event-Driven Systems
- [ ] Phase VIII — Production Engineering
- [ ] Phase IX — Distributed Systems
- [ ] Phase X — Cloud Native and Platform Engineering
- [ ] Phase XI — Go + AI Engineering
- [ ] Phase XII — Architecture, System Design and Career

---

# Definition of Done for a Module

A module is considered complete when the learner can:

- explain its main concepts in their own words;
- implement the core concepts without blindly copying code;
- recognize common mistakes and trade-offs;
- write and test idiomatic Go appropriate to the module;
- use relevant Go tooling;
- complete the required exercises/labs;
- connect the subject to a realistic engineering scenario.

---

# Repository Philosophy

This repository is both a **learning laboratory** and a **public record of engineering growth**.

Code is expected to improve over time. Early solutions should not be retroactively overengineered just to look advanced. When later modules introduce new techniques, earlier implementations may be revisited explicitly so that the trade-offs and evolution remain visible.

Mistakes, refactorings, benchmarks, debugging sessions, architecture decisions, and lessons learned are valuable parts of the journey.

---

# Contributing

This repository is intended to help other learners as well.

Contributions that improve explanations, fix mistakes, add useful exercises, improve translations, or make examples more idiomatic are welcome. Changes should preserve the curriculum's core principle:

> **Learn idiomatic Go first; introduce abstraction only when the problem justifies it.**

---

# Versioning and Evolution

The curriculum currently targets **Go 1.27**.

Go and its ecosystem will continue to evolve. Examples, tooling, dependencies, and recommendations should therefore be reviewed over time. When behavior is version-specific, it should be documented explicitly.

---

# Final Goal

By the end of this journey, the learner should not merely be able to write Go syntax. They should be able to reason about, build, test, debug, secure, observe, optimize, deploy, and evolve production software written in Go—and explain the engineering trade-offs behind their decisions.

The ultimate objective is professional readiness for **Go software engineering roles**, backed by practical projects and a public portfolio.
