# Módulo 00 — Ambiente Profissional Go 1.27

> Construindo uma base limpa e reproduzível para a Go Engineering Journey.

[English](./README.md) · [Grade Principal](../../CURRICULUM.pt-BR.md)

## Status

**Fase:** 0 — Ambiente Profissional  
**Módulo:** 00  
**Versão de referência:** Go 1.27  
**Status:** Em andamento

---

## Objetivos de Aprendizado

Ao final deste módulo, você deverá ser capaz de:

- validar uma instalação do Go;
- explicar em alto nível como código-fonte Go se transforma em um executável;
- utilizar comandos essenciais da toolchain do Go;
- inspecionar o ambiente Go;
- explicar `GOROOT`, `GOPATH`, `GOCACHE` e `GOMODCACHE`;
- compreender os papéis de `GOOS` e `GOARCH`;
- explicar a diferença entre módulo, package e arquivo-fonte;
- inicializar um Go Module;
- compilar e executar um programa Go simples;
- explicar `package main` e `func main`;
- formatar código usando as ferramentas oficiais do Go;
- realizar verificações estáticas básicas com `go vet`;
- explorar documentação usando `go doc`;
- inspecionar packages com `go list`;
- compreender os papéis básicos de `go.mod` e `go.sum`.

---

# 1. O que é Go?

Go é uma linguagem de programação compilada e estaticamente tipada, projetada com forte ênfase em simplicidade, manutenibilidade, compilação rápida, concorrência e engenharia de software prática.

Ela foi originalmente projetada no Google por Robert Griesemer, Rob Pike e Ken Thompson.

Algumas características que influenciam fortemente a maneira como software Go é escrito incluem:

- sintaxe simples;
- compilação rápida;
- garbage collection;
- primitivas de concorrência integradas à linguagem;
- composição em vez da herança tradicional de classes;
- interfaces implementadas implicitamente;
- uma Standard Library poderosa;
- ferramentas padronizadas;
- geração simples de executáveis nativos.

Uma ideia recorrente durante toda esta jornada será que aprender Go não significa apenas aprender sua sintaxe. Significa também aprender as convenções e a filosofia de engenharia do seu ecossistema.

---

# 2. Linguagem Compilada

Considere:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

Isso é código-fonte. A máquina não executa esse texto diretamente.

De maneira simplificada:

```text
Código-fonte Go
      │
      ▼
Compilador/toolchain Go
      │
      ▼
Código de máquina / executável
      │
      ▼
Sistema operacional executa
```

Esse modelo será aprofundado posteriormente quando estudarmos compiladores, linking, sistemas operacionais, memória e o runtime do Go.

Por enquanto, lembre-se:

> Go normalmente produz executáveis nativos a partir do código-fonte.

---

# 3. Verificando a Instalação do Go

Execute:

```bash
go version
```

Um resultado típico se parece com:

```text
go version go1.27.x linux/amd64
```

A saída fornece informações importantes:

```text
go version go1.27.x linux/amd64
           └──────┘ └───┘ └───┘
             versão    SO  arq.
```

A versão patch e a plataforma exatas poderão ser diferentes em sua máquina.

---

# 4. Conhecendo a Go Toolchain

Execute:

```bash
go help
```

Go fornece uma ferramenta de linha de comando unificada para grande parte das tarefas cotidianas de desenvolvimento.

Comandos importantes incluem:

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

Aprenderemos esses comandos gradualmente em vez de tentar memorizá-los todos agora.

A ideia importante é:

> A toolchain oficial faz parte da experiência de desenvolvimento da linguagem Go.

Essa padronização é uma das forças do Go.

---

# 5. Inspecionando o Ambiente

Execute:

```bash
go env
```

Esse comando apresenta a configuração do ambiente Go.

Algumas variáveis que devemos reconhecer neste momento:

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

Você ainda não precisa memorizar todas as variáveis.

Porém, deverá começar a compreender o significado dessas variáveis específicas.

---

# 6. `GOOS` e `GOARCH`

`GOOS` identifica o sistema operacional alvo.

Exemplos:

```text
linux
windows
darwin
freebsd
```

`GOARCH` identifica a arquitetura de CPU alvo.

Exemplos:

```text
amd64
arm64
386
```

Consulte os valores da sua máquina:

```bash
go env GOOS
go env GOARCH
```

Essas variáveis são particularmente importantes porque Go oferece suporte à compilação cruzada para diversos alvos.

Por exemplo, a partir de um ambiente compatível:

```bash
GOOS=linux GOARCH=amd64 go build
```

ou:

```bash
GOOS=windows GOARCH=amd64 go build
```

Estudaremos cross-compilation adequadamente mais adiante. Neste momento, a ideia principal é entender que sistema operacional e arquitetura alvo são conceitos explícitos na toolchain do Go.

---

# 7. `GOROOT`

Consulte:

```bash
go env GOROOT
```

`GOROOT` aponta para a árvore da instalação/toolchain do Go.

Conceitualmente:

```text
GOROOT
  │
  ├── toolchain do Go
  ├── fontes da Standard Library
  └── arquivos de suporte da instalação
```

Em uma instalação moderna normal, geralmente não devemos configurar `GOROOT` manualmente.

Um bom modelo mental é:

> `GOROOT` está relacionado à própria instalação do Go.

---

# 8. `GOPATH`

Consulte:

```bash
go env GOPATH
```

Historicamente, `GOPATH` possuía papel central na organização do código-fonte e das dependências Go.

Projetos Go antigos normalmente ficavam em uma estrutura semelhante a:

```text
$GOPATH/
└── src/
    └── example.com/
        └── project/
```

O desenvolvimento Go moderno utiliza **Go Modules**, portanto nossos projetos não precisam mais ficar dentro de `$GOPATH/src`.

`GOPATH` ainda existe e continua sendo utilizado pela toolchain para algumas finalidades, mas não deve ser confundido com o diretório raiz de todo projeto Go moderno.

---

# 9. `GOCACHE`

Consulte:

```bash
go env GOCACHE
```

`GOCACHE` identifica o build cache.

Quando apropriado, Go pode reutilizar artefatos de compilação gerados anteriormente em vez de reconstruir tudo do zero.

Conceitualmente:

```text
código-fonte
    │
    ▼
compilador
    │
    ├── artefatos reutilizáveis ──► GOCACHE
    │
    ▼
executável
```

Isso contribui para ciclos rápidos de desenvolvimento.

---

# 10. `GOMODCACHE`

Consulte:

```bash
go env GOMODCACHE
```

`GOMODCACHE` é o cache de módulos baixados.

Ele armazena conteúdo de módulos obtidos como dependências.

Não confunda com `GOCACHE`:

| Variável | Finalidade principal |
|---|---|
| `GOCACHE` | Cache de artefatos gerados durante builds |
| `GOMODCACHE` | Cache de módulos Go baixados |

Essa diferença ficará ainda mais clara quando começarmos a utilizar módulos externos.

---

# 11. Go Modules

Um **module** é uma coleção de packages Go relacionados e versionados em conjunto.

Projetos Go modernos normalmente utilizam um arquivo `go.mod` para definir seu módulo.

Crie um diretório:

```bash
mkdir hello-go
cd hello-go
```

Inicialize o módulo:

```bash
go mod init example.com/hello-go
```

Isso cria:

```text
hello-go/
└── go.mod
```

Um `go.mod` mínimo se parece com:

```go
module example.com/hello-go

go 1.27
```

O module path é um identificador do módulo.

Em um projeto real, ele normalmente corresponde à localização do repositório, por exemplo:

```text
github.com/username/project
```

Estudaremos module paths, versionamento semântico, resolução de dependências e os comandos de módulos com muito mais profundidade posteriormente.

---

# 12. Module vs. Package vs. Arquivo-Fonte

Esses conceitos não devem ser confundidos.

Uma hierarquia simplificada:

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

Um **arquivo-fonte** é um arquivo `.go`.

Um **package** agrupa arquivos-fonte Go que pertencem juntos.

Um **module** contém um ou mais packages versionados em conjunto.

No nosso primeiro programa:

```text
hello-go/                    ← diretório do module
├── go.mod                   ← definição do module
└── main.go                  ← arquivo-fonte Go
                              └─ pertence ao package main
```

Módulos posteriores irão refinar esse modelo mental.

---

# 13. Seu Primeiro Programa

Crie:

```text
main.go
```

com:

```go
package main

import "fmt"

func main() {
	fmt.Println("Go Professional Training")
	fmt.Println("Module 00")
	fmt.Println("Environment ready!")
}
```

O projeto agora possui:

```text
hello-go/
├── go.mod
└── main.go
```

Observe como a estrutura é pequena.

Intencionalmente, **não** estamos criando diretórios como:

```text
cmd/
internal/
pkg/
services/
repositories/
controllers/
```

Atualmente não existe nenhum problema que exija isso.

Isso é proposital.

> A estrutura deve crescer a partir de necessidades reais, não de templates copiados.

---

# 14. `package main`

Todo arquivo-fonte Go começa com uma declaração de package.

Nosso arquivo contém:

```go
package main
```

O package chamado `main` possui um papel especial: ele é usado para definir programas executáveis.

Um package `main` executável precisa de:

```go
func main()
```

Essa função é o ponto de entrada do programa.

Conceitualmente:

```text
Sistema operacional inicia o executável
                 │
                 ▼
            package main
                 │
                 ▼
              main()
                 │
                 ▼
         lógica do programa
```

---

# 15. Imports

Nosso programa contém:

```go
import "fmt"
```

`fmt` é um package da Standard Library do Go.

Utilizamos:

```go
fmt.Println(...)
```

para escrever saída formatada.

O `P` maiúsculo de `Println` possui significado.

Em Go, identificadores iniciados por letra maiúscula são exportados pelo package.

Go não utiliza palavras-chave como:

```text
public
private
protected
```

para essa finalidade.

Estudaremos visibilidade e APIs de packages profundamente mais adiante.

---

# 16. Executando o Programa

A partir do diretório do módulo:

```bash
go run .
```

Saída esperada:

```text
Go Professional Training
Module 00
Environment ready!
```

O ponto representa o package no diretório atual.

Em alto nível, `go run` compila o que for necessário e executa o programa resultante para você.

É conveniente durante o desenvolvimento.

---

# 17. Compilando o Programa

Agora execute:

```bash
go build
```

Go cria um executável para o package `main`.

No Linux/macOS, dependendo do nome do módulo/diretório, poderemos executar algo semelhante a:

```bash
./hello-go
```

No Windows, o executável gerado normalmente terá extensão `.exe`.

Saída esperada:

```text
Go Professional Training
Module 00
Environment ready!
```

A diferença conceitual é:

```text
go run .
   │
   ├── cria executável temporário
   └── executa

go build
   │
   └── cria o artefato executável
```

Essa diferença se tornará importante em fluxos profissionais de build e deployment.

---

# 18. Formatando Código Go

Go possui um padrão oficial de formatação.

Execute:

```bash
gofmt -w main.go
```

ou:

```bash
go fmt ./...
```

O primeiro comando formata diretamente o arquivo informado.

O segundo formata os packages selecionados pelo padrão `./...`.

Uma das forças culturais do Go é a padronização da formatação.

Em vez de equipes gastarem muito tempo discutindo estilo de formatação, Go fornece um formatador canônico.

> Formatação é responsabilidade da ferramenta, não gosto pessoal.

---

# 19. Verificações Estáticas Básicas com `go vet`

Execute:

```bash
go vet ./...
```

`go vet` analisa código Go em busca de construções suspeitas que possam indicar erros.

É útil distinguir conceitualmente algumas ferramentas:

```text
gofmt / go fmt
    │
    └── formatação

go vet
    │
    └── construções suspeitas / verificações estáticas

go test
    │
    └── execução de testes

Staticcheck e outros linters
    │
    └── análise estática adicional
```

Existe alguma sobreposição, mas essas ferramentas não possuem exatamente a mesma finalidade.

Estudaremos cada uma delas posteriormente.

---

# 20. Documentação pelo Terminal

A documentação Go pode ser explorada sem sair do terminal.

Experimente:

```bash
go doc fmt
```

Depois:

```bash
go doc fmt.Println
```

Esse é um hábito valioso.

Em vez de imediatamente pesquisar na internet cada API, aprenda também a consultar a documentação disponível pela própria toolchain.

---

# 21. Inspecionando Packages com `go list`

Experimente:

```bash
go list std
```

Isso lista packages da Standard Library.

Você encontrará packages como:

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

Não é esperado que você os aprenda agora.

A lição importante é que Go possui uma Standard Library substancial e fornece ferramentas para inspecioná-la.

---

# 22. `go.mod` e `go.sum`

Neste momento, nosso projeto talvez precise apenas de:

```text
go.mod
```

Quando dependências forem introduzidas, também será comum encontrar:

```text
go.sum
```

Uma distinção simplificada:

```text
go.mod
  └── declara o module e requisitos de dependências

go.sum
  └── registra checksums criptográficos usados para verificar conteúdo de módulos
```

`go.sum` não deve ser descrito simplesmente como equivalente Go de um lock file. Seu papel é diferente.

Estudaremos gerenciamento de dependências adequadamente no Módulo 1 e posteriormente nos módulos profissionais de Go.

---

# 23. Primeiro Laboratório

Crie dentro deste repositório:

```text
modules/
└── module-00/
    ├── README.md
    ├── README.pt-BR.md
    └── hello-go/
        ├── go.mod
        └── main.go
```

Dentro de `hello-go`, inicialize o módulo.

Um module path alinhado ao repositório pode ser:

```bash
go mod init github.com/blss-tico/go-engineering-journey/modules/module-00/hello-go
```

Depois crie:

```go
package main

import "fmt"

func main() {
	fmt.Println("Go Professional Training")
	fmt.Println("Module 00")
	fmt.Println("Environment ready!")
}
```

Execute a sequência completa de validação:

```bash
go fmt ./...
go vet ./...
go run .
go build
```

Depois execute o binário gerado.

### Saída esperada

```text
Go Professional Training
Module 00
Environment ready!
```

---

# 24. Exercício de Investigação do Ambiente

Utilize a toolchain para descobrir os valores da **sua máquina**:

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

Não copie valores deste documento.

O objetivo é investigar seu ambiente real de desenvolvimento.

### Meu ambiente

Preencha depois de executar os comandos:

```text
Versão do Go:
GOOS:
GOARCH:
GOROOT:
GOPATH:
GOCACHE:
GOMODCACHE:
GOTOOLCHAIN:
```

Evite versionar secrets ou caminhos sensíveis específicos da máquina caso seu ambiente contenha algum.

---

# 25. Exercícios

## Exercício 1 — Explique a hierarquia

Sem consultar a seção anterior, explique:

```text
Module → Package → Arquivo-fonte
```

Escreva com suas próprias palavras.

## Exercício 2 — Explique o ambiente

Explique a diferença entre:

```text
GOROOT
GOPATH
GOCACHE
GOMODCACHE
```

## Exercício 3 — `run` vs. `build`

Explique qual diferença prática você observa entre:

```bash
go run .
```

e:

```bash
go build
```

## Exercício 4 — Explore a Standard Library

Utilize:

```bash
go doc
```

para investigar pelo menos três packages da Standard Library.

Sugestões:

```text
fmt
os
strings
```

Escreva uma frase explicando a finalidade principal de cada package.

## Exercício 5 — Explore a toolchain

Execute:

```bash
go help
```

Escolha três comandos que ainda não estudamos profundamente e escreva uma hipótese curta sobre a finalidade de cada um.

Não há problema se a resposta não estiver perfeita. Voltaremos a eles.

---

# 26. Desafio — Observando Cross-Compilation

Ainda não trate este exercício como um processo de deployment em produção.

Primeiro consulte:

```bash
go env GOOS GOARCH
```

Depois investigue combinações de alvos disponíveis pela toolchain:

```bash
go tool dist list
```

Observe o formato:

```text
GOOS/GOARCH
```

Escolha um alvo diferente do seu ambiente atual e identifique o que mudaria no comando de build.

O objetivo é compreender o conceito, não dominar cross-compilation neste momento.

---

# 27. Checkpoint

Antes de considerar o Módulo 00 concluído, você deverá conseguir responder estas perguntas sem simplesmente repetir definições:

1. O que significa dizer que Go é uma linguagem compilada?
2. Quais informações `go version` fornece?
3. O que são `GOOS` e `GOARCH`?
4. O que é `GOROOT`?
5. Por que `GOPATH` é menos central para a estrutura dos projetos do que era historicamente?
6. Qual a diferença entre `GOCACHE` e `GOMODCACHE`?
7. O que é um Go Module?
8. O que é um Go package?
9. O que é um arquivo-fonte Go?
10. O que `go mod init` faz?
11. Por que um programa executável utiliza `package main`?
12. Qual é o papel de `func main()`?
13. Qual a diferença prática entre `go run .` e `go build`?
14. Por que `gofmt` é importante para a cultura Go?
15. O que `go vet` procura detectar?
16. Como consultar documentação pelo terminal?
17. Qual é a finalidade básica de `go.mod`?
18. Qual é a finalidade básica de `go.sum`?
19. Por que ainda não estamos criando uma estrutura complexa de diretórios?

Se alguma resposta não estiver clara, esse assunto merece revisão antes de avançarmos.

---

# 28. Notas de Engenharia

## Por que começamos com um projeto tão pequeno?

Porque arquitetura deve resolver problemas reais.

Nosso programa atualmente possui uma responsabilidade e poucas linhas de código. Adicionar várias camadas e diretórios aumentaria a carga cognitiva sem resolver um problema real de engenharia.

Posteriormente, quando nossas aplicações desenvolverem limites e complexidade reais, teremos motivos concretos para introduzir estrutura.

Isso nos permitirá observar uma progressão importante:

```text
problema simples
      │
      ▼
solução simples
      │
      ▼
novos requisitos
      │
      ▼
novas restrições
      │
      ▼
estrutura justificada
```

Essa progressão é mais educativa do que começar com um template enorme contendo abstrações que ainda não compreendemos.

## Standard Library primeiro

Durante a jornada, frequentemente investigaremos se a Standard Library do Go já oferece aquilo de que precisamos antes de adicionar dependências externas.

Isso **não** significa que bibliotecas de terceiros sejam ruins.

Significa que dependências devem ser escolhidas conscientemente.

---

# 29. O que Aprendi

Complete esta seção depois do laboratório.

### Conceitos que ficaram claros

- 
- 
- 

### Conceitos que ainda preciso revisar

- 
- 
- 

### Algo que me surpreendeu

- 

### Perguntas para a próxima aula

- 

---

# 30. Comandos Utilizados Neste Módulo

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

Não memorize essa lista mecanicamente. Aprenda qual problema cada comando resolve.

---

# 31. Critério de Conclusão

O Módulo 00 estará concluído quando:

- [ ] Go 1.27 estiver instalado e validado.
- [ ] A investigação do ambiente estiver concluída.
- [ ] `hello-go` possuir seu próprio `go.mod`.
- [ ] O programa executar corretamente com `go run .`.
- [ ] O programa compilar corretamente com `go build`.
- [ ] O executável gerado funcionar corretamente.
- [ ] `go fmt ./...` concluir sem problemas.
- [ ] `go vet ./...` concluir sem problemas.
- [ ] Os exercícios tiverem sido respondidos.
- [ ] As perguntas do checkpoint puderem ser explicadas com suas próprias palavras.
- [ ] A seção **O que Aprendi** estiver preenchida.

---

# Próximo Passo

Depois de concluir e revisar este módulo, seguiremos para:

**Módulo 01 — Go Toolchain**

Nele deixaremos de apenas utilizar alguns comandos e passaremos a compreender a toolchain de desenvolvimento Go de maneira mais sistemática.
