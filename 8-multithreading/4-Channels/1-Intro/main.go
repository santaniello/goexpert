package main

import "fmt"

// Thread 1 (Principal)
func main() {
	channel := make(chan string) // Aqui criamos um canal de string vazio

	//Thread 2
	go func() {
		channel <- "Hello World" // Aqui enviamos uma mensagem para o canal
	}()

	// Lendo a mensagem do canal na thread principal (1)
	msg := <-channel // Aqui recebemos uma mensagem do canal
	fmt.Println(msg) // Aqui imprimimos a mensagem recebida
}
