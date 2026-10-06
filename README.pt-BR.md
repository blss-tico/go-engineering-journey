# Go Engineering Journey

> Uma jornada prática dos fundamentos de Go à engenharia de software profissional.

[English](./README.md) · [Grade completa](./CURRICULUM.pt-BR.md)

## Sobre

**Go Engineering Journey** é um projeto de estudo prático e de longo prazo cujo objetivo é a formação como engenheiro de software profissional especializado em Go.

A proposta vai muito além de aprender a sintaxe da linguagem. A jornada combina Go com os conhecimentos de ciência da computação e engenharia necessários para projetar, construir, testar, depurar, proteger, observar, otimizar, implantar e evoluir software de produção.

A formação atualmente tem como referência o **Go 1.27** e foi planejada para evoluir junto com a linguagem e seu ecossistema.

Este repositório também funciona como um registro público de aprendizado: exercícios, laboratórios, experimentos, anotações, benchmarks, erros, refatorações e projetos fazem parte da jornada.

## Princípio Central

> **Aprenda primeiro Go idiomático; introduza abstrações somente quando o problema as justificar.**

Os exemplos e exercícios priorizam as convenções da comunidade Go, simplicidade e a Standard Library.

Abordagens arquiteturais como Clean Architecture, Arquitetura Hexagonal, Domain-Driven Design e microservices são intencionalmente introduzidas mais tarde, depois de construirmos a base necessária para compreender seus benefícios, custos e trade-offs.

## O que Esta Jornada Abrange

A formação progride dos fundamentos até assuntos profissionais avançados:

- Fundamentos da linguagem e Go idiomático
- Toolchain, modules, packages e gerenciamento de dependências
- Estruturas de Dados e Algoritmos
- Memória, runtime, sistemas operacionais e redes
- Testes, debugging, análise estática, profiling e performance
- Git e colaboração profissional
- SQL, PostgreSQL, NoSQL, Redis e caching
- HTTP, REST, gRPC, GraphQL, WebSockets e SSE
- Concorrência e paralelismo
- Runtime e scheduler do Go
- RabbitMQ, NATS/JetStream, Kafka e sistemas orientados a eventos
- Segurança de aplicações e engenharia de resiliência
- Observabilidade com OpenTelemetry, Prometheus e Grafana
- Sistemas distribuídos
- Docker e Kubernetes
- CI/CD, software supply chain, Platform Engineering e SRE
- Engenharia de IA com Go
- Arquitetura de Software e System Design
- Open Source, comunicação técnica e preparação para entrevistas

Consulte a **[grade completa](./CURRICULUM.pt-BR.md)** para conhecer todas as fases, módulos, projetos e checkpoints profissionais.

## Roadmap de Aprendizado

| Fase | Foco |
|---|---|
| 0 | Ambiente Profissional |
| I | Fundamentos de Go |
| II | Fundamentos de Ciência da Computação |
| III | Go Profissional |
| IV | Engenharia de Dados |
| V | Engenharia Backend |
| VI | Concorrência e Paralelismo |
| VII | Mensageria e Sistemas Orientados a Eventos |
| VIII | Engenharia de Produção |
| IX | Sistemas Distribuídos |
| X | Cloud Native e Platform Engineering |
| XI | Go + Engenharia de IA |
| XII | Arquitetura, System Design e Carreira |

A jornada completa contém **50 módulos (0–49)** e uma sequência de projetos progressivamente mais próximos de sistemas reais.

## Projetos

A prática evolui junto com a formação, em vez de ser uma coleção de tutoriais desconectados.

| # | Projeto | Foco Principal |
|---|---|---|
| 01 | CLI Profissional | Fundamentos, filesystem, errors e testes |
| 02 | Biblioteca de DSA | Estruturas de dados, algoritmos e generics |
| 03 | Servidor HTTP | `net/http`, redes e testes |
| 04 | REST API + PostgreSQL | APIs backend e persistência relacional |
| 05 | API + PostgreSQL + Redis | Cache e padrões de produção |
| 06 | Serviço gRPC | Protocol Buffers e RPC |
| 07 | Aplicação Event-Driven | Mensageria e garantias de entrega |
| 08 | Engine de Processamento Concorrente | Goroutines, channels e sincronização |
| 09 | Serviços Distribuídos | Confiabilidade e sistemas distribuídos |
| 10 | Aplicação Cloud-Native | Containers, Kubernetes e observabilidade |
| 11 | Aplicação Go + IA | Integração com LLM, RAG/tools e produção |
| 12 | Sistema Final | Engenharia de produção end-to-end |

Projetos maiores de portfólio poderão futuramente ganhar repositórios próprios. Este continuará sendo o registro central e índice da formação.

## Método de Estudo

Os assuntos normalmente seguem esta progressão:

**Teoria → Funcionamento Interno → Exemplo Mínimo → Implementação → Revisão de Go Idiomático → Erros Comuns → Exercício → Desafio → Aplicação Profissional → Revisão**

O objetivo é domínio, não velocidade. Qualquer assunto poderá ser aprofundado antes de avançarmos para o módulo seguinte.

## Progresso Atual

**Fase atual:** Fase 0 — Ambiente Profissional  
**Módulo atual:** Módulo 0 — Ambiente Profissional Go 1.27

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

O progresso será atualizado conforme a jornada avançar.

## Estrutura do Repositório

A estrutura evoluirá naturalmente conforme surgirem novas necessidades. Uma possível direção é:

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

O repositório não será superestruturado prematuramente. Diretórios e abstrações devem aparecer quando houver uma razão real para existirem.

## Filosofia de Engenharia

Este repositório preserva intencionalmente a evolução do processo de aprendizado.

É esperado que o código inicial seja mais simples que o código orientado a produção desenvolvido posteriormente. Soluções antigas não devem ser artificialmente reescritas apenas para parecerem sofisticadas.

Conforme novos conceitos forem aprendidos, soluções anteriores poderão ser revisitadas para estudarmos:

- por que determinada refatoração é útil;
- qual problema uma abstração resolve;
- qual complexidade ela introduz;
- se a performance realmente melhorou;
- se o novo design continua idiomático em Go.

Erros e experimentos que não funcionaram também são evidências úteis de engenharia quando suas lições são documentadas.

## Versão do Go

A formação atualmente utiliza como referência:

```text
Go 1.27
```

Comportamentos específicos de versão deverão ser documentados quando relevantes. O repositório poderá ser atualizado conforme novas versões de Go se tornarem apropriadas para a formação.

## Para Quem é Este Projeto?

Este repositório pode ser útil para:

- quem está aprendendo Go desde o início;
- desenvolvedores vindos de outras linguagens;
- desenvolvedores backend migrando para Go;
- engenheiros interessados em concorrência e sistemas distribuídos;
- profissionais de DevOps, SRE ou Platform Engineering que desejam aprofundar Go;
- estudantes se preparando para vagas profissionais de Go;
- qualquer pessoa procurando uma trilha estruturada que vá além da sintaxe.

## Contribuindo

Este projeto é público não apenas para documentar meu próprio aprendizado, mas também com a esperança de que possa ajudar outros estudantes.

Contribuições são bem-vindas quando melhorarem o material de estudo, por exemplo:

- corrigindo erros técnicos;
- melhorando explicações;
- melhorando traduções entre inglês e português;
- sugerindo exercícios úteis;
- corrigindo código Go não idiomático;
- melhorando testes ou documentação;
- adicionando referências relevantes.

Por favor, preserve o princípio principal do projeto:

> **Go idiomático primeiro. Abstrações quando justificadas.**

## Idiomas

- **English:** [README.md](./README.md) · [Curriculum](./CURRICULUM.md)
- **Português (Brasil):** [README.pt-BR.md](./README.pt-BR.md) · [Grade completa](./CURRICULUM.pt-BR.md)

## Licença

Este repositório é distribuído sob os termos da [MIT License](./LICENSE).

## Objetivo Final

O objetivo não é simplesmente dizer:

> “Eu sei Go.”

O objetivo é poder dizer:

> **“Eu consigo desenvolver software de produção com Go, compreendo os sistemas sobre os quais ele executa e consigo explicar os trade-offs por trás das minhas decisões.”**

Essa é a jornada.
