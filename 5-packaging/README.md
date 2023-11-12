# Packaging

O "go mod" é um recurso introduzido no Go 1.11 que trata da gestão de dependências e do controle de versões em projetos Go. Ele oferece uma maneira mais moderna e robusta de gerenciar as dependências em comparação com as abordagens anteriores, como o uso de GOPATH e a ferramenta "dep".

Aqui estão algumas características e funcionalidades-chave do "go mod":

1.  Módulos: Com o "go mod", você trabalha com módulos, que são coleções de pacotes Go que são versionados e gerenciados como uma única unidade. Um módulo é representado por um diretório raiz que contém um arquivo chamado go.mod.

2. go.mod: O arquivo go.mod é o arquivo de manifesto de um módulo. Ele lista as dependências do módulo, bem como suas versões. Você pode adicionar e atualizar dependências diretamente neste arquivo, e o Go irá garantir que as versões corretas das dependências sejam baixadas e usadas em seu projeto.

3. Versionamento semântico: O "go mod" segue o versionamento semântico, o que significa que as versões das dependências são especificadas de forma que você possa garantir a compatibilidade com seu código. Isso evita problemas de quebra de código devido a atualizações não intencionais de dependências.

4. Controle de versões: O "go mod" permite que você especifique as versões das dependências que deseja usar em seu projeto. Isso dá a você um controle preciso sobre as versões de suas dependências.

5. Compatibilidade com o GitHub e outros repositórios: O "go mod" é compatível com repositórios Git, incluindo o GitHub. Você pode importar dependências diretamente de repositórios Git, e o Go irá gerenciar essas dependências automaticamente.

6. Simplificação do ambiente de desenvolvimento: O "go mod" simplifica o ambiente de desenvolvimento, removendo a necessidade de definir GOPATH globalmente e permite que você trabalhe em projetos Go em qualquer diretório, não apenas em um GOPATH específico.

Para começar a usar o "go mod", você pode executar o seguinte comando dentro do diretório do seu projeto Go:

```bash 
go mod init nome-do-seu-modulo
```

Isso criará um arquivo go.mod inicial para o seu projeto e você pode começar a adicionar suas dependências a partir daí. O Go irá gerenciar automaticamente as dependências com base no arquivo go.mod.

## Nomenclatura de módulos em Go

O padrão para nomear módulos em Go é usar o caminho do repositório Git onde o código-fonte do módulo está hospedado. Isso geralmente significa usar o caminho completo do repositório no GitHub ou em outra plataforma de hospedagem de código.

O formato geral para nomear módulos é o seguinte:

github.com/usuario/nome-do-modulo

Onde:

- github.com é o provedor de hospedagem (pode ser substituído por outros como gitlab.com, bitbucket.org, etc.).
- usuario é o nome do usuário ou organização que possui o repositório.
- nome-do-modulo é o nome do repositório Git que contém o código-fonte do módulo.

## Baixar e atualizar dependências

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

## Trabalhando com Workspaces 

[Documentação](https://go.dev/doc/tutorial/workspaces)

Para você trabalhar com projetos multi módulo localmente sem precisar publicar esses módulos no Github, precisamo usar um recurso no da linguagem chamado Go Workspaces.

Na pasta pai (usamos como exemplo a pasta exemplo-2), iniciamos o nosso workspace infomando os módulos que vamos trabalhar:

```bash
go work init ./math ./sistema
```

Repare que um arquivo go.work foi criado na pasta exemplo-2;

Após isso, repare que conseguimos usar o nosso modulo math dentro do módulo sistema mesmo sem ele estar publicado no github.

***OBS: A ideia é que nem versionemos esse arquivo go.work pois após publicarmos esses módulos no github, ao dar um go mod tidy o go consiga encontra-los e realizar o download deles...***
