package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, cep := range os.Args[1:] {
		req := callCepApi(cep)
		defer req.Body.Close()
		data := readResponseBody(req)
		file := createFile()
		defer file.Close()
		saveCepInFile(file, data)
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("Cidade: ", data.Localidade)
	}
}

func saveCepInFile(file *os.File, data *ViaCEP) {
	/*
	 A função fmt.Sprintf na linguagem de programação Go é usada para formatar e retornar uma string, sem escrever os dados em um io.Writer ou qualquer outra saída externa. Basicamente, ela permite que você crie uma string formatada da mesma forma que as funções fmt.Printf e fmt.Fprintf, mas, em vez de imprimir ou escrever a string em algum lugar, ela simplesmente retorna a string.
	*/
	_, err := file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
	if err != nil {
		/*
		  A função fmt.Fprintf na linguagem de programação Go é usada para formatar e escrever dados em um io.Writer. Em termos simples, ela permite que você crie uma string formatada, semelhante à função fmt.Sprintf, mas em vez de retornar a string, ela escreve diretamente no io.Writer fornecido.
		  No caso os.Stderr, é a saída de erro padrão do sistema operacional.
		*/
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}
}

func createFile() *os.File {
	file, err := os.Create("cidade.txt")
	if err != nil {
		/*
		  A função fmt.Fprintf na linguagem de programação Go é usada para formatar e escrever dados em um io.Writer. Em termos simples, ela permite que você crie uma string formatada, semelhante à função fmt.Sprintf, mas em vez de retornar a string, ela escreve diretamente no io.Writer fornecido.
		  No caso os.Stderr, é a saída de erro padrão do sistema operacional.
		*/
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}
	return file
}

func readResponseBody(req *http.Response) *ViaCEP {
	res, err := io.ReadAll(req.Body)
	if err != nil {
		/*
		  A função fmt.Fprintf na linguagem de programação Go é usada para formatar e escrever dados em um io.Writer. Em termos simples, ela permite que você crie uma string formatada, semelhante à função fmt.Sprintf, mas em vez de retornar a string, ela escreve diretamente no io.Writer fornecido.
		  No caso os.Stderr, é a saída de erro padrão do sistema operacional.
		*/
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	var data ViaCEP
	err = json.Unmarshal(res, &data)
	if err != nil {
		/*
		  A função fmt.Fprintf na linguagem de programação Go é usada para formatar e escrever dados em um io.Writer. Em termos simples, ela permite que você crie uma string formatada, semelhante à função fmt.Sprintf, mas em vez de retornar a string, ela escreve diretamente no io.Writer fornecido.
		  No caso os.Stderr, é a saída de erro padrão do sistema operacional.
		*/
		fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
	}
	return &data
}

func callCepApi(cep string) *http.Response {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		/*
		  A função fmt.Fprintf na linguagem de programação Go é usada para formatar e escrever dados em um io.Writer. Em termos simples, ela permite que você crie uma string formatada, semelhante à função fmt.Sprintf, mas em vez de retornar a string, ela escreve diretamente no io.Writer fornecido.
		  No caso os.Stderr, é a saída de erro padrão do sistema operacional.
		*/
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
	}
	return req
}
