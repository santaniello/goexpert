package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

/*
*
Nesse exemplo, iremos chamar o nosso webserver que criamos na pasta /with_cancel/exemplo_webserver
e caso ele não responda em até 3 segundos, abortamos a requisição ea mensagem context deadline exceeded
será exibida no console.

Esse exemplo visa demonstrar que podemos controlar o contexto tanto do lado do cliente, quanto do servidor
e isso é muito importante pois caso uma requisição seja abortada por alguma razão, nosso servidor pode parar o processamento na hora
e não continuar processando em background
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
