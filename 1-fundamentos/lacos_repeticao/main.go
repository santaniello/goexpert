package main

import "fmt"

/*
No go temos apenas o for como laço de repetição
*/
func main() {
	// For padrão
	for i := 0; i < 10; i++ {
		println(i)
	}

	numerosSlice := []string{"Paulo", "José", "Maria"}

	// For com range
	for i, v := range numerosSlice {
		fmt.Printf("indice %d valor %s\n", i, v)
	}

	// For parecido com while

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Abaixo segue um exemplo de loop infinito utilizado para consumir uma fila e etc ..

	//for{
	//	println("Loop Infinito ")
	//}
}
