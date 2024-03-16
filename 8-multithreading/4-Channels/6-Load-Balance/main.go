package main

import "fmt"

func worker(id int, data <-chan int) {
	for value := range data {
		fmt.Printf("Worker %d received %d\n", id, value)
	}
}

func main() {
	qtdWorkers := 100      // quanto maior a quantidade de workers, mais rápido será o processamento
	qtdData := 100000      // quantidade de dados a serem processados
	data := make(chan int) // inicializa o canal

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data) // inicia a go routine para o worker i
	}

	for i := 0; i < qtdData; i++ {
		data <- i // envia o valor i para o canal que está sendo lido pelos workers
	}
}
