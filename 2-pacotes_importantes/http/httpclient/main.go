package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Chamando Get ...............")
	Get("https://www.google.com", time.Second)
	fmt.Println("Chamando POST ...............")
	Post("https://www.google.com", time.Second)
	fmt.Println("Chamando Request Customizada ...............")
	requestCustomizada("https://www.google.com", time.Second)
	fmt.Println("Chamando Request Customizada com Context ...............")
	requestCustomizadaComContext("https://www.google.com", time.Microsecond)
}

func Get(url string, duration time.Duration) {
	client := http.Client{Timeout: duration}
	req, err := client.Get(url)
	// O defer é um statement. ele atrasa a execução desse trecho de código sendo executado por ultimo.
	defer req.Body.Close()
	if err != nil {
		panic(err)
	}
	res := readResponseBody(req)
	println(string(res))
}

func Post(url string, duration time.Duration) {
	client := http.Client{Timeout: duration}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "wesley"}`))
	resp, err := client.Post(url, "application/json", jsonVar)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	// Jogando nosso ResponseBody na saida do console
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

func requestCustomizada(url string, duration time.Duration) {
	// criamos o client
	client := http.Client{Timeout: duration}
	// criamos a request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	// Aqui nós fazemos a junção do client com a request usando a função Do
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

/*
Em Go (também conhecida como Golang), o pacote context fornece uma maneira de transportar prazos, cancelamentos, e outros valores entre processos. Vamos discutir os principais conceitos e usos desse pacote.

Para que servem?

Cancelamentos: Imagine que você tenha um serviço que faz várias chamadas a outros serviços. Se um dos serviços demorar muito, talvez você queira cancelar todas as outras chamadas. O contexto permite que você propague um sinal de cancelamento para todas as goroutines que foram derivadas dele.

Timeouts e Deadlines: Junto com cancelamentos, você pode especificar um prazo (deadline) após o qual o contexto será cancelado. Isso é útil para definir um tempo máximo que um processo ou tarefa pode levar.

Transporte de Valores: O contexto permite transportar valores específicos de um processo para outro. No entanto, é importante ter cautela ao usar esta funcionalidade, pois pode tornar o código difícil de entender e manter.

*/

func requestCustomizadaComContext(url string, duration time.Duration) {
	// Retorna um contexto vazio. Geralmente é o ponto de partida para a criação de novos contextos.
	ctx := context.Background()
	/*
		 Criando um contexto com timeout que poderá ser cancelado após o tempo de timeout ou através
		da função de cancelamento que estamos retornando
	*/
	ctx, cancel := context.WithTimeout(ctx, duration)
	/*
		O uso do defer cancel() é uma prática comum e importante. Mesmo que o contexto seja cancelado
		automaticamente após o timeout, é uma boa prática sempre chamar cancel() para liberar os recursos
		associados a ele o mais rápido possível. Usando defer, garantimos que a função de cancelamento
		será chamada assim que a função principal (main() neste caso) retornar, garantindo que os
		recursos sejam liberados.
	*/
	defer cancel()

	// Aqui passamos nosso contexto que tem a mesma função que o http.Client{Timeout: duration} que por de baixo dos panos acaba criando um contexto...
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

func readResponseBody(req *http.Response) []byte {
	// Aqui estamos fazendo a leitura do Body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	return res
}
