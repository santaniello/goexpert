package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Toda a requisição tem um contexto atrelado a ela, aqui estamos recuperando esse contexto
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		// Se o usuário não cancelar a request, após 5 segundos ele irá cair aqui
		// Imprime no comand line stdout
		log.Println("Request processada com sucesso")
		// Imprime no browser
		w.Write([]byte("Request processada com sucesso"))
		// Se por alguma razão o usuário cancelar a request com um ctrl + c ele vai cair aqui
	case <-ctx.Done():
		// Imprime no comand line stdout
		log.Println("Request cancelada pelo cliente")
	}
}
