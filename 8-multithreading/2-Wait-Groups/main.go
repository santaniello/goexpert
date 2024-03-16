package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done() // cada vez que passar por aqui, ele irá decrementar 1 das nossas 25 operações ...
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25) // Aqui adicionamos o total de operação que vamos usar 10 da task A + 10 da task B + 5 da função anônima
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)
	// Thread 4 - Podemos criar uma go routine a partir de uma função anônima
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done() // cada vez que passar por aqui, ele irá decrementar 1 das nossas 25 operações ...
		}
	}()
	// Ele irá esperar as 25 operações serem executadas para encerrar o programa. OBS: se houvesse comando a seguir, ele continuaria a execução ...
	waitGroup.Wait()
}
