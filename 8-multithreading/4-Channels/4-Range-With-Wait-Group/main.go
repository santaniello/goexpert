package main

import (
	"fmt"
	"sync"
)

/*
A ideia aqui é basicamente criarmos 2 for/ranges sincronizados ... um para escrever e outro para ler
A cada iteração do writer, ele envia um valor para o canal, bloqueando o mesmo e o reader recebe esse mesmo valor e imprime na tela.
Quando o canal for desbloqueado, o writer envia o próximo valor e o reader recebe e imprime na tela, e assim por diante.
*/
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)
	go writer(ch)
	go reader(ch, &wg) // Nesse caso, ao contrário do exemplo 3-Reader, estamos utilizando um waitgroup para sincronizar a execução do reader com o writer. Isso fará com que o reader só encerre sua execução quando o writer tiver enviado todos os 10 valores para o canal e não fará com que o programa encerre sua execução abruptamente.
	wg.Wait()
}

// Func reader lê 10 valores do canal ch e depois encerra sua execução
func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}

// Func writer escreve 10 valores no canal ch e depois fecha o canal
func writer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Aqui indicamos que não haverá mais nenhum valor a ser enviado para o canal, logo, o reader irá encerrar sua execução. Isso fará com que não seja gerardo um deadlock pois o reader não irá ficar esperando um novo valor indefinidamente.
}
