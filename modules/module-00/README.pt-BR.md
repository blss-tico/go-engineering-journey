# Módulo 00 — Ambiente Profissional de Desenvolvimento Go

> **Go Engineering Journey** · Fase 0 — Ambiente Profissional · Go 1.27

[English](README.md) · [Grade em português](../../CURRICULUM.pt-BR.md) · [Página inicial](../../README.pt-BR.md)

## Objetivo

Construir um ambiente de desenvolvimento Go confiável e reproduzível, compreendendo como a toolchain, os módulos, os packages, as dependências, os executáveis compilados e o Git trabalham em conjunto. Este módulo é uma base prática — não é uma aplicação de produção nem um exercício de arquitetura de software.

**Pré-requisitos:** terminal, Git e uma toolchain compatível com Go 1.27. Os exemplos foram praticados com Go 1.27.1 em Linux/amd64. Adapte comandos e caminhos ao seu sistema operacional.

## Objetivos de aprendizagem

Ao concluir o módulo, você deverá conseguir:

1. Inspecionar a toolchain instalada e as principais variáveis de ambiente do Go.
2. Explicar as diferenças entre repositórios, módulos, packages e arquivos-fonte.
3. Executar, formatar, inspecionar, testar, analisar e compilar um programa Go simples.
4. Fazer cross-compilation e inspecionar o destino e os metadados de um binário.
5. Criar e importar um package dentro do mesmo módulo.
6. Adicionar um módulo externo, compreender `go.mod` e `go.sum` e diferenciar grafos de módulos e de packages.
7. Utilizar Git para revisar, registrar, sincronizar e documentar um laboratório reproduzível.

## Estrutura do repositório

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
    └── bin/               # Executáveis locais; ignorados pelo Git
```

O diretório `bin/` é gerado localmente e não precisa existir no repositório. O laboratório contém intencionalmente um módulo e dois packages.

## 1. Inspecionando a instalação do Go

```bash
go version
go env GOROOT GOPATH GOCACHE GOMODCACHE GOTOOLCHAIN
go env GOOS GOARCH CGO_ENABLED
```

| Configuração | Significado |
| --- | --- |
| `GOROOT` | Local da distribuição Go e da biblioteca padrão |
| `GOPATH` | Diretório de trabalho que inclui, por padrão, a pasta superior do cache de módulos |
| `GOCACHE` | Cache de compilação |
| `GOMODCACHE` | Cache dos módulos baixados |
| `GOTOOLCHAIN` | Comportamento de seleção da toolchain |
| `GOOS`, `GOARCH` | Sistema operacional e arquitetura de destino |
| `CGO_ENABLED` | Indica se cgo está habilitado na compilação |

**Reflexão:** Por que não devemos guardar o repositório dentro de `GOROOT`? Qual é a diferença entre `GOMODCACHE` e o diretório do projeto?

## 2. Compreendendo os limites do módulo

A partir da raiz do repositório:

```bash
cd modules/module-00/hello-go
cat go.mod
go env GOMOD
go list -m
go list -f 'Name={{.Name}} ImportPath={{.ImportPath}}' .
```

O module path é:

```text
github.com/blss-tico/go-engineering-journey/modules/module-00/hello-go
```

Um **repositório** armazena arquivos versionados. Um **módulo** é uma unidade versionada identificada por `go.mod`. Um **package** reúne arquivos Go compilados em conjunto dentro de um diretório. Um **arquivo-fonte** é um arquivo `.go`. Um módulo pode conter vários packages; o import path é uma identidade lógica, não um caminho físico específico da máquina.

## 3. Executando, formatando, analisando e compilando

```bash
gofmt -w main.go greeting/greeting.go
go run .
go vet ./...
go test ./...
mkdir -p bin
go build -o bin/hello-linux-amd64 .
./bin/hello-linux-amd64
```

Em Linux/amd64, o programa deverá exibir uma saudação, o destino `linux/amd64` e uma frase em português com capitalização de título. `go test ./...` pode informar `[no test files]`; isso não comprova cobertura de testes funcionais.

`go run .` compila e executa temporariamente o programa; `go build -o ... .` grava um executável no caminho escolhido. `gofmt` é o formatador canônico do Go. `go vet` realiza verificações estáticas selecionadas, mas não prova que o programa esteja correto.

## 4. Inspecionando binários e fazendo cross-compilation

```bash
file bin/hello-linux-amd64
go version -m bin/hello-linux-amd64
GOOS=freebsd GOARCH=amd64 CGO_ENABLED=0 go build -o bin/hello-freebsd-amd64 .
file bin/hello-freebsd-amd64
go version -m bin/hello-freebsd-amd64
```

`GOOS` e `GOARCH` definem o **destino da compilação**, que não precisa coincidir com o sistema operacional e a arquitetura da máquina de desenvolvimento. Um executável FreeBSD compilado no Linux geralmente não pode ser executado diretamente no Linux. Compilações cruzadas com cgo podem exigir um compilador cruzado apropriado; por isso o exemplo desabilita cgo explicitamente.

`go version -m` pode apresentar module path, opções de compilação e informações de revisão Git. Quando existem metadados de VCS, `vcs.modified=true` indica que a toolchain detectou alterações locais na compilação; `false` indica estado limpo segundo essa verificação. Esses metadados ajudam na rastreabilidade, mas não constituem uma atestação completa da cadeia de suprimentos.

## 5. Criando um package local

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

O identificador `Message`, iniciado com letra maiúscula, é **exportado** e pode ser acessado por outro package. Um identificador `message`, com inicial minúscula, não seria acessível a partir de `main`.

O package local `greeting` **não** precisa de outro `go.mod`: seu import path é formado pelo module path principal acrescido de `/greeting`. Um `go.mod` aninhado criaria outro limite de módulo.

Investigue a resolução de packages:

```bash
go list ./...
go list -f 'Name={{.Name}} ImportPath={{.ImportPath}}' ./greeting
go list -f '{{.Imports}}' .
go list -f '{{.Module.Path}}' ./greeting
go list -f '{{.ImportPath}} {{.Standard}}' fmt
```

## 6. Adicionando e mantendo um módulo externo

O exemplo importa dois packages, `cases` e `language`, de **um único módulo externo**, `golang.org/x/text`.

```bash
go get golang.org/x/text@latest
go mod tidy
cat go.mod
cat go.sum
go list -m all
go mod verify
```

**Dependência observada no laboratório:** `golang.org/x/text v0.42.0`. Esse valor registra o exercício; não significa que todos devam instalar sempre essa versão. `@latest` pode resolver para versões diferentes ao longo do tempo e alterar a versão mínima de Go do módulo. Revise as mudanças em `go.mod` antes de fazer commit.

- **`go.mod`** declara o módulo principal, a versão mínima de Go e os requisitos de módulos. Go seleciona versões usando *Minimal Version Selection* (MVS).
- **`go.sum`** registra checksums de conteúdo e metadados de módulos. **Não** é um lockfile tradicional de versões.
- **`// indirect`** indica um requisito que não fornece package diretamente importado pelo módulo principal (além de alguns casos relacionados ao grafo de módulos).
- **`go get`** altera requisitos; **`go mod tidy`** reconcilia requisitos e checksums com o código e os testes relevantes.
- **`go mod verify`** verifica a integridade dos módulos baixados no cache; não audita vulnerabilidades nem executa os testes do programa.

### Grafo de módulos versus grafo de packages

```bash
go mod graph
go list -deps -f '{{if not .Standard}}{{.ImportPath}}{{end}}' . | sed '/^$/d'
go list -f '{{.Imports}}' .
```

Cada comando responde a uma pergunta diferente:

| Comando | O que apresenta |
| --- | --- |
| `go mod graph` | Relações de requisitos entre módulos |
| `go list -m all` | Módulos selecionados na lista de compilação |
| `go list -deps ...` | Packages necessários para compilar o package selecionado |
| `go list -f '{{.Imports}}' .` | Imports diretos do package principal |

No exercício registrado, `x/text v0.42.0` declarou requisitos de `x/tools v0.49.0`, `x/mod v0.41.0` e `x/sync v0.23.0` no **grafo de módulos**. Os packages desses módulos não apareceram na listagem filtrada de dependências do executável. Um módulo presente na lista de compilação não necessariamente tem código compilado dentro do binário.

Diretórios chamados `internal` impõem restrições de importação: código fora da árvore permitida pelo diretório pai não pode importar diretamente esses packages.

## 7. Fluxo de trabalho com Git e rastreabilidade

Mantenha os binários locais fora dos commits adicionando `bin/` ao `.gitignore` do repositório.

A partir da raiz do repositório:

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

Este é um fluxo ilustrativo; **não repita o commit se o laboratório já foi registrado**. `git diff` não exibe arquivos não rastreados até que sejam preparados. `git status` distingue alterações do working tree e do índice.

Se o push for rejeitado porque o remoto avançou, investigue antes de integrar:

```bash
git fetch origin
git status
git log --oneline --graph --decorate -12 --all
git log --oneline HEAD..origin/main
git log --oneline origin/main..HEAD
```

Com working tree limpo e commits locais ainda não publicados, `git rebase origin/main` pode ser apropriado. Resolva conflitos conscientemente; não use `git push --force` como resposta automática a uma rejeição.

## 8. Exercícios práticos

1. Explique `GOROOT`, `GOPATH`, `GOCACHE` e `GOMODCACHE` sem consultar anotações.
2. Compare `go run .` e `go build -o bin/hello .`.
3. Compile para Linux e FreeBSD e inspecione os binários com `file` e `go version -m`.
4. Explique por que `greeting` é resolvido localmente sem outra declaração `require`.
5. Altere temporariamente `Message` para `message`. Preveja e explique o erro de compilação; depois restaure o código.
6. Compare `go mod graph`, `go list -m all` e `go list -deps`.
7. Explique por que `go.sum` não é um lockfile e por que `// indirect` não significa “não instalado”.
8. Inspecione o histórico Git e confirme a sincronização entre branches local e remota.

## 9. Critérios de conclusão — Definition of Done

- [ ] Go 1.27 está instalado e as variáveis de ambiente foram compreendidas.
- [ ] `hello-go` possui um módulo e pelo menos dois packages.
- [ ] `go run .` funciona e sua saída é compreendida.
- [ ] Formatação, `go vet` e `go test ./...` terminam sem erros.
- [ ] Os destinos de compilação Linux e FreeBSD foram inspecionados.
- [ ] Sei explicar `go.mod`, `go.sum`, requisitos diretos/indiretos e checksums.
- [ ] Comparei os grafos de módulos e packages.
- [ ] `go mod verify` termina com sucesso.
- [ ] Os binários são ignorados e as alterações de código foram registradas.
- [ ] O repositório está sincronizado e `git status` está limpo.
- [ ] Completei a reflexão abaixo com minhas próprias palavras.

## 10. O que aprendi

Complete esta seção depois do laboratório:

- **Conceito mais importante:** …
- **Observação mais surpreendente:** …
- **Um problema que encontrei e como resolvi:** …
- **O que quero explorar mais:** …

## Próximo módulo

**Módulo 01 — Go Toolchain Avançada:** aprofundamento no comportamento do compilador e linker, comandos da toolchain, documentação, diagnósticos e fluxos de desenvolvimento reproduzíveis.

---

**Referências oficiais:** [Documentação Go](https://go.dev/doc/) · [Referência de módulos](https://go.dev/ref/mod) · [Documentação do comando Go](https://pkg.go.dev/cmd/go) · [Effective Go](https://go.dev/doc/effective_go)
