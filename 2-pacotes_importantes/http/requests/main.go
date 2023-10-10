package main

import (
	"io"
	"net/http"
)

func main() {
	req := call("https://www.google.com")
	// O defer é um statement. ele atrasa a execução desse trecho de código sendo executado por ultimo.
	defer req.Body.Close()
	res := readResponseBody(req)
	println(string(res))

}

func call(url string) *http.Response {
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return req
}

func readResponseBody(req *http.Response) []byte {
	// Aqui estamos fazendo a leitura do Body
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	return res
}
