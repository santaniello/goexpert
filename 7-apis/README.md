## GO project layout

Segue um repo que é um modelo interessante de estrutura que podemos seguir para nossoas projetos em Golang:

https://github.com/golang-standards/project-layout

## Pastas importantes

- internal: Essa pasta contém os arquivos que são exclusivos do nosso negócio e não queremos compartilhar através do go mod.
- pkg: Ao contrário da pasta internal, a pasta pkg contém libraries que nós podemos compartilhar com outros projetos.
- cmd/nome_da_aplicacao_folder: Diretório onde vai ficar o seu executável. Aqui fica localizado o arquivo main.go
- configs: contém os arquivos para o bootstrap da aplicação
- test: Aplicativos de teste externos adicionais e dados de teste.
- api: Guarda as especificações da nossa api (swagger).

## Middleware


Um middleware em Go (Golang) é um termo usado para descrever uma função ou um conjunto de funções que são executadas antes ou depois de um manipulador HTTP. Em outras palavras, eles são intermediários que processam as solicitações HTTP antes de chegarem ao manipulador final (ou endpoint), e/ou processam as respostas antes de serem enviadas de volta ao cliente.

As principais características e usos de um middleware em Go são:

- Processamento Intermediário: Middlewares são usados para executar algum tipo de processamento nas solicitações ou respostas. Por exemplo, eles podem ser usados para adicionar cabeçalhos comuns, registrar solicitações, autenticar usuários, manipular erros, etc.

- Cadeia de Manipuladores: Em uma aplicação web, os middlewares são normalmente organizados em uma cadeia. Cada middleware pode decidir se passa a solicitação para o próximo middleware na cadeia ou termina o processamento.

- Modificação de Solicitações e Respostas: Middlewares têm a capacidade de modificar a solicitação e a resposta. Por exemplo, um middleware pode adicionar informações adicionais ao contexto da solicitação, que podem ser usadas por manipuladores subsequentes.

- Reusabilidade: Um middleware pode ser reutilizado em diferentes partes de uma aplicação. Por exemplo, um middleware de logging pode ser usado em todas as rotas para registrar informações sobre as solicitações.

- Sintaxe em Go: Em Go, um middleware é normalmente implementado como uma função que recebe e retorna um http.Handler. Aqui está um exemplo simplificado:

```go
func MyMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Código do middleware aqui

        next.ServeHTTP(w, r)
    })
}
```

Este middleware recebe um manipulador next (o próximo na cadeia) e retorna um novo manipulador. Dentro da função, você pode executar o código necessário antes e/ou depois de chamar next.ServeHTTP(w, r).

Em resumo, middlewares em Go são uma ferramenta poderosa para adicionar funcionalidades transversais (como logging, autenticação, e manipulação de erros) de maneira modular e reutilizável em aplicações web.

**OBS: Um middleware em Go é basicamente uma função que recebe e retorna um http.Handler.** 

### Como registro um middleware usando um multiplexer router como o Chi ou Gorila mux ?

```go
r.Use(MyCustomMiddleware)
r.Use(LoggerMiddleware)
r.Use(MyCustomMiddleware)
http.ListenAndServe(":8080", r)
```

Eles serão executados na ordem em que são registrados e para compartilharmos valores entre eles, precisamos adicionar esses valores no contexto da request para que ele esteja acessivel a outros 
middlewares.



## JWT

[Documentação](https://jwt.io/)

JWT, que significa JSON Web Token, é um padrão aberto (RFC 7519) para criar tokens de acesso que permitem a transmissão de informações entre duas partes de forma segura e compacta. É amplamente utilizado na autenticação e autorização em aplicações web e móveis. Aqui estão os pontos principais sobre o que é JWT e como ele funciona:

### Estrutura do JWT

Um JWT é composto por três partes, separadas por pontos (.):

Exemplo:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

1. Header (Cabeçalho): O cabeçalho tipicamente consiste em duas partes: o tipo do token, que é JWT, e o algoritmo de assinatura utilizado, como HMAC SHA256 ou RSA.

2. Payload (Carga Útil): O payload contém as afirmações ou reivindicações. Estas reivindicações podem ser:

   - Reivindicações Registradas: Um conjunto de reivindicações predefinidas que não são obrigatórias, mas recomendadas, como iss (emissor), exp (tempo de expiração), sub (assunto do token), etc.
   - Reivindicações Públicas: Estas podem ser definidas à vontade pelos produtores e consumidores do JWT. Porém, devem ser usadas com cuidado para evitar colisões.
   - Reivindicações Privadas: Estas são reivindicações criadas para compartilhar informações entre partes que concordam com o uso delas e não são registradas nem públicas.

Signature (Assinatura): Para criar a assinatura, o cabeçalho codificado em base64, o payload codificado em base64 e uma senha são combinados e então assinados com o algoritmo especificado no cabeçalho.

### Funcionamento

- Autenticação: Na autenticação, após o usuário se logar com suas credenciais, o servidor cria um JWT com informações do usuário e uma assinatura segura. O servidor então retorna esse JWT ao cliente. O cliente armazena esse token e o envia junto às requisições subsequentes para o servidor.
- Autorização: Quando uma aplicação cliente envia uma requisição ao servidor, ela inclui o JWT, geralmente no cabeçalho Authorization. O servidor, por sua vez, valida o JWT e, se válido, processa a requisição.
- Segurança: Como a assinatura é criada usando o cabeçalho e o payload, é possível verificar se o conteúdo do token foi alterado. Se alguém tentar alterar o payload ou o cabeçalho sem a chave secreta, a verificação da assinatura falhará no servidor.

### Vantagens e Desvantagens

**Vantagens:**

- Compacto: Pode ser enviado através de URLs, parâmetros POST ou no cabeçalho HTTP.
- Autocontido: Carrega todas as informações necessárias, o que evita a necessidade de consultar o banco de dados mais de uma vez.
- Padronizado: Facilita a interoperabilidade entre diferentes sistemas.

**Desvantagens:**

- Armazenamento e Segurança: Como é armazenado no lado do cliente, o armazenamento seguro do JWT é crítico.
- Inflexibilidade: Uma vez emitido, não pode ser revogado antes da expiração, a menos que mecanismos adicionais sejam implementados.

**Conclusão**

JWTs são uma maneira eficaz de transmitir informações de forma segura entre partes e são amplamente usados em sistemas modernos para autenticação e autorização. No entanto, requerem um manuseio cuidadoso, especialmente no que diz respeito à segurança e gerenciamento de tokens.


## Swagger

- [Documentação](https://github.com/swaggo/swag)
- [OpenAPI](https://www.openapis.org/)

### Qual a diferença entre Swagger e OpenAPI

1. Swagger: Originalmente, Swagger foi a primeira encarnação da especificação para descrever APIs RESTful. Ele incluía uma especificação e um conjunto de ferramentas para ajudar no desenvolvimento de APIs, como a interface do usuário do Swagger para visualizar documentação de API, o Swagger Editor para escrever especificações de API e o Swagger Codegen para gerar código cliente e servidor a partir da especificação da API.

2. OpenAPI: Em 2015, a especificação Swagger foi doada à Linux Foundation e passou a ser gerida pelo OpenAPI Initiative, um projeto colaborativo que envolve várias empresas. A partir desse ponto, a especificação Swagger foi renomeada para OpenAPI Specification (OAS). O OpenAPI é uma especificação mais ampla e robusta para descrever APIs RESTful. A OpenAPI Specification evoluiu com mais recursos e funcionalidades em comparação com a versão original do Swagger.

**Resumindo:**

Swagger agora se refere mais comumente às ferramentas desenvolvidas pela SmartBear Software que trabalham com a especificação OpenAPI, incluindo Swagger UI, Swagger Editor e Swagger Codegen.
OpenAPI Specification (OAS) é o nome atual da especificação anteriormente conhecida como Swagger. Ela define um formato padrão para descrever APIs RESTful.
Portanto, enquanto Swagger começou como a especificação, agora é mais sobre as ferramentas que suportam a OpenAPI Specification, que é o padrão atual para descrever APIs.

### Instalação

Primeiro, devemos instalar o binario do swago

```bash
go install github.com/swaggo/swag/cmd/swag@latest 
```

Após isso, precisamos garantir que a variavel de ambiente GOPATH esteja exportada nos arquivos ~/.zshrc se vc estiver usando zshell ou ~/.bashrc se você estiver usando bash

Exemplo de configuração:

```shell
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

teste o comando swag no terminal e veja se possui alguma saida.

### Gerando e configurando a doc

Exemplo de documentação de uma api:

```go
// GetProducts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Success      200       {array}   entity.Product
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {}
```

Precisamos agora utilizar um comando via terminal para gerar essa doc:

```bash
swag init -g cmd/apis/main.go 
```

***Esse comando deve ser executado toda vez que documentarmos uma nova api para que as docs na pasta docs sejam alteradas e a documentação apareça.***

***Importante:*** 
- A opção -g indica a localização do arquivo main.go que é o arquivo base que será utilizado pelo swag para a geração das docs.
  Se o arquivo main estivesse na raiz da app, não seria necessário a utilização desse parâmetro.
  
- Repare que uma pasta docs foi criada na raiz do projeto com as docs do mesmo.
- No arquivo main, fizemos a configuração do swagger e para isso existem 2 imports muitos importantes pois um deles indica onde está a pasta docs comas documentações dp projeto. Sem esses imports, não funcionará:
  - _ "github.com/santaniello/apis/docs"
  - httpSwagger "github.com/swaggo/http-swagger"

Agora, por ultimo precisamos configurar no arquivo main o endpoint onde acessaremos o swagger:

```go
r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
```





