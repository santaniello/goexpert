# GO (Golang)

## O que á linguagem GO (Golang)?

- Linguagem de programação open source que tem o objetivo de tronar os programadores mais produtivos;
- Expressiva, concisa, limpa e eficiente;
- Foi criada para aproveitar ao máximo dos recursos multicore e de rede;
- Rápida compilação e ao mesmo tempo trabaçha com garbage collection;
- Rápida, estaticamente tipada, compilada mas que ao memo tempo parece até uma linguagem dinamicamente tirada e interpetada;
- Compilada em apenas um arquivo binário;
- Retrocompativel com versões antigas da linguagem;


## O que á linguagem GO não é 

- Uma linguagem de programação que resolverá todos os seus problemas;
- Não é uma linguagem dinâmica;
- Não é uma linguagem interpretada;
- Não é uma linguagem com muitos recursos / firulas;

## Motivação para a criação do  GO

- Limitação de algumas principais linguagens usadas na Google como Python, Java e C++;
- Python: problemas com lentidão;
- C/C++: Muita complexidade e demorado para compilar;
- Java: Complexidade gerada ao longo do tempo / verbosidade da linguagem;
- Multithreading e Concorrência: Não nasceram natovamente pensando nisso;
- Simplicidade;
- Framework de testes e profiling nativos;
- Detecção de Race Conditions;
- Deploy absurdamente simples;
- Baixa curva de aprendizado;

## Instalando o Go

![Instalação do GO](https://go.dev/)


### Quais são as principais variáveis de ambiente do GO

```go
go env
```

Exemplo de saída:

```bash
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/fsantaniello/.cache/go-build"
GOENV="/home/fsantaniello/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/fsantaniello/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/fsantaniello/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.20.5"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-O2 -g"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-O2 -g"
CGO_FFLAGS="-O2 -g"
CGO_LDFLAGS="-O2 -g"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build3740808225=/tmp/go-build -gno-record-gcc-switches"
```

#### GOPATH

GOPATH é uma variável de ambiente usada pelo Go (linguagem de programação) até a versão Go 1.10 
para determinar o local do workspace de trabalho do Go. A partir da introdução dos Go Modules
na versão Go 1.11, o GOPATH não é mais essencial para a maioria dos projetos, mas ainda pode ser
útil em alguns contextos.

O GOPATH contém três subdiretórios principais:

- src: Onde o código-fonte dos projetos Go é mantido. Por exemplo, se você tem um projeto em github.com/meuprojeto, então o código-fonte estaria em $GOPATH/src/github.com/meuprojeto.

- pkg: Este diretório contém arquivos compilados (.a) de bibliotecas.

- bin: Contém os binários de programas e ferramentas compilados.

Quando você instala um pacote ou ferramenta Go usando go get, o código-fonte é baixado para o diretório src, e se houver algum binário associado, ele é compilado e colocado no diretório bin.

Antes da introdução dos Go Modules, o GOPATH era fundamental para gerenciar dependências e organizar seu código. No entanto, com os Go Modules, é possível criar e trabalhar com projetos Go fora do diretório GOPATH, tornando o desenvolvimento mais flexível.

Apesar dessa mudança, muitas ferramentas e sistemas antigos ainda podem depender da configuração do GOPATH. Se você estiver trabalhando com código ou ferramentas mais antigas, ou em ambientes específicos que ainda esperam a existência de um GOPATH, ainda precisará configurá-lo.


#### GOMODCACHE

GOMODCACHE é uma variável de ambiente no sistema Go (a linguagem de programação Go, muitas vezes referida apenas como "Go" ou "Golang").

A partir do Go 1.15, GOMODCACHE define o diretório base usado para armazenar os arquivos de cache de módulos baixados. Se essa variável de ambiente não estiver definida, o padrão será $GOPATH/pkg/mod. Esses arquivos de cache são essenciais para o sistema de módulos do Go, pois permitem que as dependências de um projeto sejam buscadas, armazenadas em cache e reutilizadas entre compilações e projetos diferentes.

Por exemplo, se você usa Go modules (introduzido em Go 1.11) e executa go get para baixar uma dependência, essa dependência é armazenada no diretório GOMODCACHE. Isso ajuda a acelerar as construções futuras e garante a reproducibilidade, pois você não precisa baixar a mesma dependência repetidamente da fonte original.

Definir ou alterar o valor de GOMODCACHE pode ser útil se você quiser armazenar o cache de módulos em um local diferente, por qualquer razão, como questões de espaço em disco ou organização de diretório.


## O que é o Go Mod ?

go mod é uma subcomando do comando go relacionado ao sistema de módulos da linguagem de programação Go (Golang). Introduzido em Go 1.11, o sistema de módulos fornece uma maneira integrada para versionar e gerenciar as dependências dos pacotes Go.

Antes da introdução dos módulos, o Go usava o GOPATH como uma maneira de organizar e gerenciar o código. No entanto, GOPATH tinha suas limitações, principalmente em relação ao gerenciamento de versões de pacotes. O sistema de módulos foi introduzido para resolver muitos desses problemas e simplificar o gerenciamento de dependências em Go.

Aqui está uma visão geral de algumas operações comuns que você pode realizar com go mod:

1. Iniciar um novo módulo:

```go
go mod init <module-name>
```
Isso criará um novo arquivo go.mod no diretório atual, que descreve o módulo, suas dependências e outras informações.

2. Adicionar dependências:

Ao usar comandos como go get, as dependências são automaticamente adicionadas ao seu arquivo go.mod.

Abaixo, pegamos a ultima versão de uma dependência:

```go
go get example.com/some/dependency@latest
``

Abaixo, pegamos uma versão específica de uma dependência:

```go
go get example.com/some/dependency@v1.2.3
```

3. Atualizar dependências:

```go
go get -u
```

4. Atualizar Todas as Dependências (Incluindo Indiretas):

```go
go get -u all
```

Este comando atualizará todas as dependências listadas em seu go.mod para suas versões mais recentes.

5. Verificar dependências removendo as que não são utilizadas e adicionando as que estão faltando:

```go
go mod tidy
```

Este comando remove quaisquer dependências não utilizadas de go.mod e atualiza o arquivo go.sum, que contém as somas de verificação esperadas para o conteúdo de cada dependência.

6. Visualizar dependências e versões:

```go
go list -m all
```

Mostra todas as dependências do módulo atual e suas versões.

7. Baixar dependências para o cache local:

```go
go mod download
```
O sistema de módulos Go, com o comando go mod, representa uma melhoria significativa na forma como as dependências são gerenciadas em Go. Permite um gerenciamento de dependências mais preciso, reprodutível e flexível em comparação com o antigo sistema GOPATH.

## Compilando projetos

Como funciona o processo de compilação no Go ?

**********************************************************
Código Go (Runtime) + Seu Código Go = Binário (Executável)
**********************************************************
Basicamente, quando compilamos nosso projeto com o comando build, ele une
o nosso código ao Runtime do Go gerando um unico arquivo binário.

Como podemos compilar o nosso projeto ?

Dentro do nosso modulo onde está o arquivo main.go, executamos o seguinte comando:

```go
go build main.go
```

e será gerado um arquivo main onde podemos fazer a execução da seguinte maneira:

```bash
./main
```

***OBS: Em um projeto que GOMOD, não precisamos especificar o nome do arquivo main,
isso fará com que o binário seja gerado com o nome do seu módulo ao invés do nome main***

Exemplo:

```go 
go build
```
Vai gerar um binário com o nome do seu módulo.

Para executar:

```bash
./meu-modulo
```

### Compilando sua aplicação Go para diferentes sistemas operacionais

Em go, ao compilar uma aplicação e gerar nosso binário, podemos falar para qual SO e arquitetura de processador podemos compilar nosso projeto.

Exemplo:

```go
GOOS=windows GOARCH=amd64 go build main.go
```

No exemplo acima, estamos gerando um binário para o windows na arquitetura amd64.

Como saber quais sistemas operacionais e arquitetura eu posso gerar meu binário ?

para isso, execute o comando abaixo para listar o sos e arquiteturas disponiveis:

```go
go tool dist list
```

OBS: As variaveis GOOS GOARCH são duas variaveis de ambiente do GO e para lista-las, podemos executar os seguintes comandos:

```go
go env  // esse comando lista todas as variaveis de ambiente do Go
```
ou

```go
go env GOOS GOARCH // lista somente as 2 que queremos
```
e a saida será: 

```bash
linux
amd64
```

pois esse é o meu SO e a arquitetura do meu processador.

[Tutorial Building Go Applications for Different Operating Systems and Architectures](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)
