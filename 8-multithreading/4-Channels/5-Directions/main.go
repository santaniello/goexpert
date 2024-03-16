package main

import "fmt"

// ch chan<- string significa que o canal ch só pode ser escrito (receive only - Canal só recebe dados)
func enviaMensagem(mensagem string, ch chan<- string) {
	ch <- mensagem
}

// ch <-chan string significa que o canal ch só pode ser lido (Send only - Canal só envia dados)
func recebeMensagem(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string)
	go enviaMensagem("Hello World", ch)
	recebeMensagem(ch)

}
