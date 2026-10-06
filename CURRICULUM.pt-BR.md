# Go Engineering Journey — Grade da Formação

> Dos fundamentos de Go à engenharia de software profissional.

**Idioma:** [English](./CURRICULUM.md) | Português (Brasil)

## Sobre esta formação

Este repositório documenta uma jornada prática e de longo prazo para a formação de um engenheiro de software profissional especializado em Go. O objetivo não é aprender apenas a sintaxe da linguagem, mas desenvolver conhecimentos de ciência da computação, engenharia de software, backend, infraestrutura, sistemas distribuídos, segurança, observabilidade, cloud native e engenharia de IA necessários no trabalho real.

A formação tem como referência o **Go 1.27** e deverá evoluir junto com a linguagem e seu ecossistema.

## Princípios da formação

1. **Go idiomático primeiro.** Exemplos, exercícios e projetos iniciais seguem as convenções da comunidade Go e favorecem soluções simples e idiomáticas.
2. **Standard Library primeiro.** Aprender o que Go já oferece antes de introduzir frameworks ou abstrações externas.
3. **Entender antes de abstrair.** Clean Architecture, Arquitetura Hexagonal, DDD, repositories, services e abordagens semelhantes entram somente depois da base necessária para avaliar seus trade-offs.
4. **Teoria + implementação.** Conceitos importantes são estudados teoricamente e depois implementados ou investigados em Go.
5. **Mentalidade de produção.** Testes, debugging, segurança, performance, observabilidade, tratamento de falhas e operação fazem parte da engenharia.
6. **Otimização baseada em evidências.** Medir antes de otimizar.
7. **Aprender em público.** Exercícios, laboratórios, documentação, histórico Git e projetos formam um registro visível da evolução e um portfólio.
8. **IA como ferramenta, não substituto do entendimento.** Desenvolvimento assistido por IA é utilizado criticamente, com revisão e verificação humana.

## Trilha de aprendizado

| Fase | Foco | Resultado |
|---|---|---|
| 0 | Ambiente Profissional | Ambiente produtivo de desenvolvimento Go |
| I | Fundamentos de Go | Domínio sólido da linguagem |
| II | Ciência da Computação | DSA, memória, SO e redes |
| III | Go Profissional | Testes, debugging, qualidade, performance, Git |
| IV | Engenharia de Dados | SQL, PostgreSQL, NoSQL, Redis |
| V | Engenharia Backend | HTTP, REST, gRPC, GraphQL e APIs |
| VI | Concorrência e Paralelismo | Modelo de concorrência e runtime do Go |
| VII | Mensageria e Event-Driven | Brokers, entrega e padrões orientados a eventos |
| VIII | Engenharia de Produção | Segurança, resiliência, observabilidade, performance |
| IX | Sistemas Distribuídos | Consistência, confiabilidade e falhas |
| X | Cloud Native e Platform Engineering | Containers, Kubernetes e CI/CD |
| XI | Engenharia de IA | Sistemas com IA construídos em Go |
| XII | Arquitetura, System Design e Carreira | Design avançado e preparação profissional |

---

## Fase 0 — Ambiente Profissional

### Módulo 0 — Ambiente Profissional Go 1.27
Instalação e validação; Linux; `GOROOT`, `GOPATH`, `GOCACHE`, `GOMODCACHE`; `GOOS`, `GOARCH`; modules; editor/IDE; terminal; repositório; primeiro build e execução.

### Módulo 1 — Go Toolchain
`go run`, `go build`, `go install`, `go test`, `go fmt`/`gofmt`, `go vet`, `go doc`, `go env`, `go list`, `go generate`, `go tool` e comandos de módulos/dependências.

## Fase I — Fundamentos de Go

### Módulo 2 — Fundamentos da Linguagem
Packages; variáveis; constantes; tipos built-in; operadores; controle de fluxo; zero values; escopo.

### Módulo 3 — Sistema de Tipos
Tipagem estática; tipos definidos e aliases; conversões; inferência; comparabilidade; identidade de tipos.

### Módulo 4 — Arrays, Slices e Maps
Arrays; internals de slices; tamanho/capacidade; `append`; `copy`; maps; `make`; armadilhas comuns.

### Módulo 5 — Strings, Bytes, Runes e Unicode
Strings; UTF-8; `byte`; `rune`; Unicode; `strings`; `bytes`.

### Módulo 6 — Structs, Pointers e Fundamentos de Memória
Structs; ponteiros; semântica de valores; `new`; `make`; modelo mental inicial de stack/heap.

### Módulo 7 — Funções, Métodos, Interfaces e Composição
Funções; múltiplos retornos; variádicas; closures; `defer`; métodos/receivers; method sets; interfaces; implementação implícita; embedding e composição.

### Módulo 8 — Errors e Generics
Tratamento idiomático; wrapping; `errors.Is`; `errors.As`; custom errors; `panic`/`recover`; type parameters; constraints; inferência; estruturas genéricas; recursos relevantes do Go 1.27.

## Fase II — Fundamentos de Ciência da Computação

### Módulo 9 — Estruturas de Dados
Arrays dinâmicos; listas ligadas; stacks; queues/deques; hash tables; sets; heaps/priority queues; árvores; BSTs; conceitos de árvores balanceadas; tries; grafos.

### Módulo 10 — Algoritmos
Complexidade; Big-O; busca; ordenação; recursão; divide and conquer; greedy; backtracking; programação dinâmica; BFS/DFS; shortest paths; ordenação topológica; problemas de entrevista.

### Módulo 11 — Memória e Fundamentos do Runtime
Stack/heap; escape analysis; alocações; GC; layout; alinhamento; locality; pressão de alocação.

### Módulo 12 — Sistemas Operacionais para Engenheiros Go
Processos/threads; scheduling; context switch; memória virtual; syscalls; file descriptors; signals; filesystems; I/O bloqueante/não bloqueante; conceitos de `epoll`; namespaces; cgroups.

### Módulo 13 — Fundamentos de Redes
TCP/IP; IP; TCP/UDP; DNS; sockets; TLS; HTTP/1.1; HTTP/2; conceitos de HTTP/3; keep-alive; connection pooling; timeouts.

## Fase III — Go Profissional

### Módulo 14 — Packages, Modules e Design de APIs
Coesão; APIs públicas/internas; `internal`; limites de módulos; dependências; compatibilidade; versionamento; ciclos; package design idiomático.

### Módulo 15 — Testes
`testing`; table-driven tests; subtests; helpers; test doubles; integração/E2E/contrato; fuzzing; property-based testing; cobertura; race detector.

### Módulo 16 — Debugging
Delve; breakpoints; breakpoints condicionais; stack traces; inspeção de goroutines; debugging concorrente; race detector; conceitos de debugging remoto.

### Módulo 17 — Qualidade e Análise Estática
`gofmt`; `go vet`; Staticcheck; linters; code review; higiene de dependências; code smells; Go idiomático.

### Módulo 18 — Benchmarking, Profiling e Performance
Benchmarks; perfis de CPU/heap/alocação; `pprof`; blocking/mutex profiles; goroutine leaks; flame graphs; PGO; metodologia de performance.

### Módulo 19 — Git Profissional e Colaboração
Fundamentos/internals; branches; merge/rebase; conflitos; commits; PRs; code review; tags; SemVer; releases; fluxos colaborativos.

## Fase IV — Engenharia de Dados

### Módulo 20 — Teoria de Bancos de Dados
Modelo relacional; modelagem; normalização/desnormalização; chaves/constraints; ACID; transações; isolamento; MVCC; locks; deadlocks; índices; B-trees; query planning.

### Módulo 21 — PostgreSQL com Go
SQL; `database/sql`; `pgx`; pools; prepared statements; transações; migrations; `EXPLAIN`; otimização; errors; testes.

### Módulo 22 — NoSQL e Escolha de Armazenamento
Key-value; documentos; wide-column; séries temporais; MongoDB como estudo; SQL vs. NoSQL; escolha orientada a requisitos.

### Módulo 23 — Redis e Cache
Redis; cache-aside; TTL; eviction; invalidation; cache distribuído; stampede; penetration; serialização; quando não usar cache.

## Fase V — Engenharia Backend

### Módulo 24 — HTTP em Profundidade
Semântica; métodos; status; headers; cookies; content negotiation; compression; conexões; timeouts; `net/http`; servidores/clientes.

### Módulo 25 — APIs REST
Modelagem de recursos; routing; middleware; validação; errors; paginação; filtros; ordenação; versionamento; OpenAPI; documentação; testes.

### Módulo 26 — Engenharia Avançada de APIs
Idempotência; rate limiting; compatibilidade; API Gateway; webhooks; WebSockets; SSE; streaming; ciclo de vida.

### Módulo 27 — gRPC e Protocol Buffers
Protobuf; serviços; unary; server/client/bidirectional streaming; interceptors; deadlines; cancellation; errors; evolução de API.

### Módulo 28 — GraphQL
Schemas; queries; mutations; resolvers; DataLoader; N+1; autorização; subscriptions; trade-offs.

**Comparação:** REST vs. gRPC vs. GraphQL vs. WebSockets vs. SSE.

## Fase VI — Concorrência e Paralelismo

### Módulo 29 — Concorrência em Go
Goroutines; channels; buffering; `select`; Mutex/RWMutex; atomics; WaitGroup; `context`; cancelamento; races; deadlocks.

### Módulo 30 — Padrões de Concorrência
Worker pools; pipelines; fan-out/fan-in; producer/consumer; semáforos; rate limiting; bounded concurrency; graceful shutdown; backpressure.

### Módulo 31 — Runtime e Scheduler do Go
Internals de goroutines; G/M/P; scheduler; work stealing; preemption; network poller; syscalls bloqueantes; interação com GC.

### Módulo 32 — Paralelismo e Performance de Baixo Nível
Concorrência vs. paralelismo; CPU-bound/I/O-bound; `GOMAXPROCS`; multicore; contention; atomics; false sharing; caches de CPU; escalabilidade.

## Fase VII — Mensageria e Sistemas Orientados a Eventos

### Módulo 33 — Teoria de Mensageria
Queues; topics; producers/consumers; pub/sub; ACK; ordering; garantias de entrega; retries; DLQ; idempotência; backpressure; evolução de schemas.

### Módulo 34 — RabbitMQ, NATS/JetStream e Kafka
Arquitetura; clientes Go; entrega/persistência; consumer groups; streams; operação; escolha tecnológica; Transactional Outbox; Saga; CQRS; conceitos de Event Sourcing.

## Fase VIII — Engenharia de Produção

### Módulo 35 — Segurança de Aplicações
Threat modeling; OWASP; TLS/mTLS; autenticação/autorização; hashing; JWT; OAuth 2.0/OIDC; secrets; validação; SQL injection; SSRF; rate limiting; segurança de dependências/supply chain; secure defaults.

### Módulo 36 — Engenharia de Resiliência
Timeouts; deadlines; retries; exponential backoff; jitter; circuit breakers; bulkheads; load shedding; backpressure; graceful degradation/shutdown.

### Módulo 37 — Observabilidade
Logs estruturados; métricas; tracing distribuído; OpenTelemetry; Prometheus; Grafana; RED; USE; SLIs/SLOs; alertas; propagação de contexto.

### Módulo 38 — Engenharia de Performance em Produção
Load/stress/soak tests; throughput; latency; percentis; capacity planning; profiling em produção; gargalos; regressões.

## Fase IX — Sistemas Distribuídos

### Módulo 39 — Fundamentos de Sistemas Distribuídos
Modelos de falha; partições; CAP; consistência; replicação; particionamento; tempo lógico; consenso; leader election; consistência eventual; transações distribuídas.

### Módulo 40 — Padrões de Confiabilidade Distribuída
Idempotência; deduplicação; distributed locks; Saga; Transactional Outbox; retries seguros; circuit breakers; recuperação; falhas parciais; mensagens duplicadas/atrasadas.

## Fase X — Cloud Native e Platform Engineering

### Módulo 41 — Containers e Docker
Containers; images; Dockerfiles; multi-stage/minimal images; networking; volumes; limites; signals; segurança; graceful shutdown; health checks.

### Módulo 42 — Kubernetes
Pods; Deployments; Services; ConfigMaps; Secrets; probes; requests/limits; HPA; Ingress; Helm; controllers/operators; Go no ecossistema Kubernetes.

### Módulo 43 — CI/CD e Software Supply Chain
CI/CD; GitHub Actions; testes/análise estática; builds; registries; SBOM; vulnerability scanning; signing; releases; estratégias de deployment; conceitos de Platform Engineering e SRE.

## Fase XI — Go + Engenharia de IA

### Módulo 44 — Engenharia de IA para Desenvolvedores Go
Fundamentos de LLMs; tokens/contexto; APIs de modelos; streaming; structured outputs; embeddings; vector search; RAG; tool/function calling; agents; avaliação; guardrails; prompt injection; observabilidade; rate limits; custos; serviços Go com IA; desenvolvimento assistido por IA; revisão de código Go gerado.

## Fase XII — Arquitetura, System Design e Carreira

### Módulo 45 — Arquitetura de Software
Somente após uma base forte em Go: monólito modular; microservices; Clean Architecture; Hexagonal/Ports and Adapters; DDD; Vertical Slice; event-driven; boundaries; trade-offs; overengineering; adaptação dos padrões ao Go idiomático.

### Módulo 46 — System Design
Requisitos funcionais/não funcionais; estimativas; throughput; latency; bandwidth; storage; load balancing; caching; partitioning; replication; fault tolerance; disponibilidade; escalabilidade; trade-offs; entrevistas.

### Módulo 47 — Open Source e Leitura de Go Real
Codebases grandes; fonte da Standard Library; projetos Go consolidados; issues; PRs; contribuições; RFCs/design docs; contribuição open source.

### Módulo 48 — Documentação Técnica e Comunicação
READMEs; documentação de packages/APIs; ADRs; RFCs; diagramas; escrita técnica; inglês técnico; comunicação de trade-offs.

### Módulo 49 — Carreira e Entrevistas
Currículo Go; portfólio GitHub; LinkedIn; entrevistas de Go/DSA/SQL/concorrência/debugging/System Design; live coding; comunicação comportamental; identificação de gaps.

---

# Projetos

| # | Projeto | Foco principal |
|---|---|---|
| 01 | CLI Profissional | Fundamentos, filesystem, errors, testes |
| 02 | Biblioteca de DSA | Estruturas, algoritmos, generics |
| 03 | Servidor HTTP | `net/http`, redes, testes |
| 04 | REST API + PostgreSQL | API e persistência relacional |
| 05 | API + PostgreSQL + Redis | Cache e padrões de produção |
| 06 | Serviço gRPC | Protobuf e RPC |
| 07 | Aplicação Event-Driven | Mensageria e garantias de entrega |
| 08 | Engine de Processamento Concorrente | Goroutines, channels, sincronização |
| 09 | Serviços Distribuídos | Confiabilidade e sistemas distribuídos |
| 10 | Aplicação Cloud-Native | Containers, Kubernetes, observabilidade |
| 11 | Aplicação Go + IA | LLM, RAG/tools e produção |
| 12 | Sistema Final | Engenharia de produção end-to-end |

Projetos maiores de portfólio poderão ganhar repositórios próprios, enquanto este repositório preserva o histórico de aprendizado, laboratórios, notas e links.

# Checkpoints Profissionais

Os checkpoints medem competências técnicas, não cargo ou anos de experiência.

### Checkpoint 1 — Fundamentos de Go
Meta: base sólida de Go para nível inicial.

### Checkpoint 2 — Go + HTTP + SQL + Git
Meta: preparação para backend Go júnior.

### Checkpoint 3 — Concorrência + Bancos + APIs + Testes
Meta: amplitude técnica de júnior forte / início de pleno.

### Checkpoint 4 — Mensageria + Observabilidade + Docker + Performance
Meta: competências de engenharia de produção de nível pleno.

### Checkpoint 5 — Sistemas Distribuídos + Kubernetes + System Design
Meta: competências técnicas avançadas associadas ao nível sênior.

### Checkpoint 6 — Runtime + Performance + Arquitetura + Platform + Open Source
Meta: profundidade avançada em engenharia Go.

---

# Método de Estudo

Cada assunto deverá normalmente seguir esta progressão:

1. **Teoria** — compreender o conceito e seu vocabulário.
2. **Funcionamento interno** — entender o que Go ou o sistema subjacente está fazendo.
3. **Exemplo mínimo** — isolar a ideia no menor programa útil possível.
4. **Implementação prática** — escrever o código.
5. **Revisão de Go idiomático** — comparar a solução com as convenções da comunidade Go.
6. **Erros comuns** — conhecer modos de falha e abordagens enganosas.
7. **Exercício** — consolidar o conceito.
8. **Desafio** — resolver um problema com menos orientação.
9. **Aplicação profissional** — relacionar o assunto com software real.
10. **Revisão** — explicar o conceito sem depender do exemplo.

O objetivo não é avançar rapidamente pelos módulos. Qualquer assunto poderá ser aprofundado antes de seguirmos adiante.

---

# Acompanhamento de Progresso

Use os seguintes marcadores:

- `[ ]` Não iniciado
- `[~]` Em andamento
- `[x]` Concluído
- `[R]` Precisa de revisão

## Progresso por fase

- [ ] Fase 0 — Ambiente Profissional
- [ ] Fase I — Fundamentos de Go
- [ ] Fase II — Fundamentos de Ciência da Computação
- [ ] Fase III — Go Profissional
- [ ] Fase IV — Engenharia de Dados
- [ ] Fase V — Engenharia Backend
- [ ] Fase VI — Concorrência e Paralelismo
- [ ] Fase VII — Mensageria e Sistemas Orientados a Eventos
- [ ] Fase VIII — Engenharia de Produção
- [ ] Fase IX — Sistemas Distribuídos
- [ ] Fase X — Cloud Native e Platform Engineering
- [ ] Fase XI — Go + Engenharia de IA
- [ ] Fase XII — Arquitetura, System Design e Carreira

---

# Critério de Conclusão de um Módulo

Um módulo será considerado concluído quando o estudante conseguir:

- explicar os principais conceitos com suas próprias palavras;
- implementar os conceitos centrais sem simplesmente copiar código;
- reconhecer erros comuns e trade-offs;
- escrever e testar Go idiomático adequado ao conteúdo estudado;
- utilizar as ferramentas Go relevantes;
- concluir os exercícios e laboratórios obrigatórios;
- relacionar o conteúdo a um cenário realista de engenharia.

---

# Filosofia do Repositório

Este repositório é simultaneamente um **laboratório de aprendizado** e um **registro público da evolução como engenheiro de software**.

É esperado que a qualidade do código evolua ao longo do tempo. Soluções iniciais não devem ser artificialmente reescritas com arquiteturas avançadas apenas para parecerem sofisticadas. Quando módulos posteriores introduzirem novas técnicas, implementações anteriores poderão ser revisitadas explicitamente, permitindo observar os trade-offs e a evolução.

Erros, refatorações, benchmarks, sessões de debugging, decisões arquiteturais e lições aprendidas também fazem parte da jornada.

---

# Contribuindo

Este repositório também pretende ajudar outros estudantes.

Contribuições que melhorem explicações, corrijam erros, adicionem exercícios úteis, aprimorem traduções ou tornem os exemplos mais idiomáticos são bem-vindas. As mudanças devem preservar o princípio central da formação:

> **Aprenda primeiro Go idiomático; introduza abstrações somente quando o problema as justificar.**

---

# Versionamento e Evolução

A formação atualmente tem como referência o **Go 1.27**.

Go e seu ecossistema continuarão evoluindo. Exemplos, ferramentas, dependências e recomendações deverão ser revisados ao longo do tempo. Quando algum comportamento depender de uma versão específica, isso deverá ser documentado explicitamente.

---

# Objetivo Final

Ao final desta jornada, o estudante não deverá apenas saber escrever a sintaxe de Go. Deverá ser capaz de raciocinar sobre, construir, testar, depurar, proteger, observar, otimizar, implantar e evoluir software de produção escrito em Go — e explicar os trade-offs de engenharia por trás de suas decisões.

O objetivo final é estar tecnicamente preparado para buscar **vagas profissionais de desenvolvimento Go**, apoiado por projetos práticos e um portfólio público.
