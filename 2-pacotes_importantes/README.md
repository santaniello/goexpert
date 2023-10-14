# Pacotes Importantes

## net/http


O pacote net/http da linguagem de programação Go (ou GoLang) é uma das bibliotecas padrão mais importantes para a criação de servidores HTTP e clientes HTTP. Ele fornece um conjunto rico de funcionalidades para trabalhar com o protocolo HTTP.

1. **Servidor HTTP:** Ele permite criar servidores HTTP de forma fácil. Com apenas algumas linhas de código, você pode ter um servidor rodando:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Olá, mundo!")
    })
    http.ListenAndServe(":8080", nil)
}
```

2. **Cliente HTTP:** O pacote também fornece funções para fazer requisições HTTP, como GET, POST e outras:

```go
resp, err := http.Get("https://www.example.com")
```

3. **FileServer:**  Você pode facilmente criar um FileServer, que é um servidor HTTP para servir arquivos estáticos, como imagens, arquivos CSS, JavaScript, etc.

```go
func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

No exemplo acima, ele irá servir tudo que está no diretório /public que no caso é a página index.html.

outro exemplo:

```go
func main() {
	// Crie um manipulador de arquivos que sirva arquivos estáticos a partir de um diretório chamado "static"
	fileServer := http.FileServer(http.Dir("static"))

	// Use o manipulador de arquivos para servir todos os pedidos que começam com "/files/"
	http.Handle("/files/", http.StripPrefix("/files/", fileServer))

	// Inicie o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Neste exemplo, se você tiver um arquivo chamado image.jpg no diretório static, você poderá acessá-lo no navegador em http://localhost:8080/files/image.jpg.

**OBS: Não temos o diretório static nesse projeto, esse exemplo foi tirado do ChatGPT !**


### ServerMux (Multiplexer) também frequentemente chamado de mux ou router

No contexto de servidores web, um multiplexer de solicitações é uma ferramenta que distribui as solicitações HTTP recebidas para seus respectivos manipuladores com base em critérios como o método HTTP (GET, POST, etc.) e o caminho da URL. Isso permite que seu aplicativo responda de maneira diferente a diferentes URLs ou diferentes tipos de solicitações.

Por exemplo, você pode querer que uma solicitação para "/home" retorne uma página inicial, enquanto uma solicitação para "/about" retorne uma página "Sobre nós". O multiplexer ajuda a dirigir cada uma dessas solicitações para a função correta que pode gerar a resposta adequada.

#### Usando DefaultServeMux

Quando você passa nil como o segundo argumento para http.ListenAndServe, o Go usa o DefaultServeMux como o multiplexer de solicitações. Em outras palavras, você está dizendo para o servidor usar o multiplexer padrão fornecido pelo pacote net/http.

Abaixo, segue um exemplo de um código que usa o mux default do Go:

```go
func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)}
```

Repare que como segundo parâmetro da função http.ListenAndServe estamos passando nil, sendo assim, o go irá usaro mux default.

#### Usando um Mux Personalizado:

Quando você passa uma instância de http.NewServeMux() para http.ListenAndServe, você está indicando que o servidor deve usar esse multiplexer personalizado que você criou, em vez do DefaultServeMux. Isso fornece a você mais controle sobre a configuração e o comportamento do seu servidor.

Em Go, o pacote net/http fornece uma implementação básica de um multiplexer chamado ServeMux. Aqui está um exemplo de como você pode usar ServeMux para definir diferentes manipuladores para diferentes caminhos:

```go
package main

import (
    "fmt"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Página inicial")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Sobre nós")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/home", homePage)
    mux.HandleFunc("/about", aboutPage)

	// Aqui passamos um mux personalizado
    http.ListenAndServe(":8080", mux)
}
```

Neste exemplo, criamos um novo ServeMux e configuramos para que as solicitações para "/home" sejam tratadas pela função homePage e solicitações para "/about" sejam tratadas pela função aboutPage. Em seguida, passamos esse mux para http.ListenAndServe para começar a ouvir e servir as solicitações.

E qual a vantagem de eu criar um mux personalizado ao invés de usar o default ?

Usar um ServeMux personalizado, ao invés do DefaultServeMux, oferece várias vantagens:

1. **Encapsulamento:** Criar seu próprio multiplexer permite um maior grau de encapsulamento. Você pode ter múltiplos multiplexers em diferentes partes do seu programa, cada um configurado de forma diferente, o que é útil se você estiver construindo bibliotecas ou módulos que precisem de suas próprias rotas.
2. **Flexibilidade:** Ao criar seu próprio ServeMux, você tem a flexibilidade de adicionar ou remover rotas em tempo de execução, se necessário. Isto pode não ser comum, mas é algo que você não poderia fazer facilmente com o DefaultServeMux sem afetar todas as outras rotas já registradas nele.
3. **Middleware e Chain Handlers:** Com multiplexers personalizados, especialmente aqueles fornecidos por frameworks de terceiros (como Gorilla Mux, Chi, etc.), você pode facilmente adicionar middlewares (funcionalidades que são executadas antes ou depois dos seus handlers principais). Isso é útil para funcionalidades como logging, autenticação, CORS e outros aspectos de tratamento de solicitação/resposta.
4. **Funcionalidades Avançadas:** Enquanto o ServeMux padrão do Go é bastante simples e realiza correspondência de prefixo para rotas, multiplexers de terceiros, como o mencionado Gorilla Mux, oferecem recursos avançados, como correspondência de padrão, extração de variáveis do caminho, restrições de método HTTP e muito mais.
5. **Isolamento de Erros:** Ao usar seu próprio multiplexer, um erro ou problema relacionado à configuração de rotas não afetará o resto do sistema que pode estar usando o DefaultServeMux.
6. **Clareza e Intenção:** Em projetos maiores, usar um multiplexer personalizado pode tornar o código mais legível, mostrando claramente qual multiplexer está sendo configurado e usado, ao invés de confiar em um estado global oculto.

Dito isso, para aplicações simples ou para começar rapidamente, o DefaultServeMux é frequentemente suficiente. À medida que as necessidades da sua aplicação se tornam mais complexas, pode fazer sentido considerar a adoção de um multiplexer personalizado ou uma solução de terceiros mais robusta.

## encoding/json

O pacote encoding/json do Go é uma biblioteca padrão usada para codificar e decodificar dados no formato JSON (JavaScript Object Notation). Ele oferece funcionalidades para converter entre estruturas Go (como structs e slices) e representações JSON, e vice-versa.

Aqui estão as principais funcionalidades e como você pode usá-las:

1. **Marshal:** Converte uma estrutura Go em uma representação JSON.

```go
type Pessoa struct {
    Nome  string `json:"nome"`
    Idade int    `json:"idade"`
}

func main() {
    p := Pessoa{Nome: "Alice", Idade: 30}
    jsonData, err := json.Marshal(p)
    if err != nil {
        log.Fatalf("Erro ao codificar: %s", err)
    }
    fmt.Println(string(jsonData))  // {"nome":"Alice","idade":30}
}
```

2. **Unmarshal:** Converte uma representação JSON em uma estrutura Go.

```go
var p Pessoa
inputData := []byte(`{"nome":"Bob","idade":25}`)
err := json.Unmarshal(inputData, &p)
if err != nil {
    log.Fatalf("Erro ao decodificar: %s", err)
}
fmt.Println(p)  // {Bob 25}
```

3. **Encoder:** Fornece uma maneira de codificar JSON diretamente em um fluxo de escrita, como um io.Writer.

```go
enc := json.NewEncoder(os.Stdout)
enc.Encode(p)
```

4. **Decoder:** Similarmente, permite decodificar JSON diretamente de um fluxo de leitura, como um io.Reader.

```go
dec := json.NewDecoder(os.Stdin)
err := dec.Decode(&p)
```

O pacote encoding/json é essencial para muitos aplicativos Go, especialmente aqueles que lidam com APIs RESTful, armazenamento de dados e qualquer outra situação onde a serialização e deserialização de dados em JSON é necessária. Ele é projetado para ser simples de usar, mas também flexível o suficiente para lidar com casos mais complexos e personalizados.

## text/template ou html/template

Em Go, templates são uma forma de gerar saídas baseadas em um modelo e dados fornecidos. O pacote text/template e o pacote html/template da biblioteca padrão do Go são usados para esse propósito.

Eles são muito similares ao JSP e a expression language do Java.

São muito pouco usados pois hoje em dia o front-end é separado do back-end.

Os exemplos estão na pasta template.

[Documentação](https://pkg.go.dev/html/template)