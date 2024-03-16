package main

import "fmt"

/*
A ideia aqui é basicamente criarmos 2 for/ranges sincronizados ... um para escrever e outro para ler
A cada iteração do writer, ele envia um valor para o canal, bloqueando o mesmo e o reader recebe esse mesmo valor e imprime na tela.
Quando o canal for desbloqueado, o writer envia o próximo valor e o reader recebe e imprime na tela, e assim por diante.
*/
func main() {
	ch := make(chan int)
	go writer(ch)
	reader(ch) // Nesse caso, não foi utilizado uma go routine para o reader, pois o writer já está em uma goroutine, logo, o reader não precisa ser executado em uma goroutine pois se fosse, ele seria executado em paralelo com o writer, e o programa encerraria sua execução.
}

// Func reader lê 10 valores do canal ch e depois encerra sua execução
func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

// Func writer escreve 10 valores no canal ch e depois fecha o canal
func writer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Aqui indicamos que não haverá mais nenhum valor a ser enviado para o canal, logo, o reader irá encerrar sua execução. Isso fará com que não seja gerardo um deadlock pois o reader não irá ficar esperando um novo valor indefinidamente.
}
