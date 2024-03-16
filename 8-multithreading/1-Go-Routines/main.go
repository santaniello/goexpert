package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")
	// Thread 4 - Podemos criar uma go routine a partir de uma função anônima
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()
	// Como não tem nada aqui, o programa vai finalizar e as threads vão ser encerradas, por isso não vemos nada no console. Para contornar isso, vamos colocar um time.Sleep() para que o programa espere 15 segundos antes de finalizar e possamos ver o resultado
	time.Sleep(15 * time.Second)
}
