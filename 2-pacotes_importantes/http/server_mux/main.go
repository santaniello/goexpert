package main

import "net/http"

/*
No contexto de servidores web, um multiplexer de solicitações é uma ferramenta que distribui as solicitações HTTP recebidas para seus respectivos manipuladores com base em critérios como o método HTTP (GET, POST, etc.) e o caminho da URL. Isso permite que seu aplicativo responda de maneira diferente a diferentes URLs ou diferentes tipos de solicitações.

Abaixo, diferentemente da pasta busca_cep_rest_api, estamos usando um mux personalizado

Usar um ServeMux personalizado, ao invés do DefaultServeMux, oferece várias vantagens:

Encapsulamento: Criar seu próprio multiplexer permite um maior grau de encapsulamento. Você pode ter múltiplos multiplexers em diferentes partes do seu programa, cada um configurado de forma diferente, o que é útil se você estiver construindo bibliotecas ou módulos que precisem de suas próprias rotas.

Flexibilidade: Ao criar seu próprio ServeMux, você tem a flexibilidade de adicionar ou remover rotas em tempo de execução, se necessário. Isto pode não ser comum, mas é algo que você não poderia fazer facilmente com o DefaultServeMux sem afetar todas as outras rotas já registradas nele.

Middleware e Chain Handlers: Com multiplexers personalizados, especialmente aqueles fornecidos por frameworks de terceiros (como Gorilla Mux, Chi, etc.), você pode facilmente adicionar middlewares (funcionalidades que são executadas antes ou depois dos seus handlers principais). Isso é útil para funcionalidades como logging, autenticação, CORS e outros aspectos de tratamento de solicitação/resposta.

Funcionalidades Avançadas: Enquanto o ServeMux padrão do Go é bastante simples e realiza correspondência de prefixo para rotas, multiplexers de terceiros, como o mencionado Gorilla Mux, oferecem recursos avançados, como correspondência de padrão, extração de variáveis do caminho, restrições de método HTTP e muito mais.

Isolamento de Erros: Ao usar seu próprio multiplexer, um erro ou problema relacionado à configuração de rotas não afetará o resto do sistema que pode estar usando o DefaultServeMux.

Clareza e Intenção: Em projetos maiores, usar um multiplexer personalizado pode tornar o código mais legível, mostrando claramente qual multiplexer está sendo configurado e usado, ao invés de confiar em um estado global oculto.

Multiplos ServerMux: Podemos criar diversos serverMux que irão atender servers diferentes inclusive simultaneamente como no exemplo deste arquivo !

Dito isso, para aplicações simples ou para começar rapidamente, o DefaultServeMux é frequentemente suficiente. À medida que as necessidades da sua aplicação se tornam mais complexas, pode fazer sentido considerar a adoção de um multiplexer personalizado ou uma solução de terceiros mais robusta.
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello World 8080"))
	})

	// Usando a struct Blog que irá lidar com a nossa Request ao invés de usar uma função anônima
	mux.Handle("/blog", Blog{title: "Teste"})
	/*
			A função http.ListenAndServe é bloqueante portanto, para subir 2 servidores simultaneamente, precisamos
			iniciar o primeiro servidor em uma goroutine separada, enquanto o segundo servidor é iniciado no thread principal.
		    Assim, ambos os servidores podem rodar simultaneamente.
	*/
	go func() {
		http.ListenAndServe(":8080", mux)
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello World 8081"))
	})
	http.ListenAndServe(":8081", mux2)
}

/*
Podemos criar uma struct que irá implementar a interface Handler através da função ServeHTTP e irá lidar com a
nossa requisição ao invés de criar uma função anônima.
*/
type Blog struct {
	title string
}

// Implementacao da Interface Handler
func (b Blog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(b.title))
}
