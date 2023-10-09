package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*****************************************************************
	   Trabalhando com Arrays
	******************************************************************/

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println("O valor de meu array no indice 1 é " + strconv.Itoa(meuArray[1]))
	fmt.Println("O tamanho de meu array é " + strconv.Itoa(len(meuArray)))
	fmt.Println("O valor do ultimo indice do meu array é " + strconv.Itoa(obterUltimoElementoArray(meuArray)))

	percorrerArray(meuArray)

	/*****************************************************************
	   Trabalhando com Slices
	******************************************************************/

	// Outra maneira de declarar e já inicializar um slice
	slice2 := []int{1, 2, 3, 4, 5}

	// Aqui imprimos o tamanho do slice a sua capacidade além dos elementos do próprio
	fmt.Printf("len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)

	// Aqui nós "fatiamos" o nosso slice criando um sub-slice vazia não incluindo o indice 0, Como resultado, você obterá uma sub-slice vazia de slice2.
	// Repare que a capacidade não será alterada
	fmt.Printf("len=%d cap=%d %v\n", len(slice2[:0]), cap(slice2[:0]), slice2[:0])

	// Aqui nós "fatiamos" o nosso slice criando um sub-slice com os 3 primeiros elementos a partir do slice2.
	// Repare que a capacidade não será alterada
	fmt.Printf("len=%d cap=%d %v\n", len(slice2[:3]), cap(slice2[:3]), slice2[:3])

	// Aqui nós "fatiamos" o nosso slice porém nós ignoramos os 3 primeiros valores pois agora temos 3: e não :3. Isso nos resultará apenas os 2 ultimos valores.
	// Repare que aqui a capacidade foi alterada pois nós ignoramos o começo
	fmt.Printf("len=%d cap=%d %v\n", len(slice2[3:]), cap(slice2[3:]), slice2[3:])

	// Adicionando um novo elemento a um Slice
	// Repare que aqui a capacidade foi dobrada pois a operação de redimensionamento de um slice é muito custosa e ele pega o tamnho atual multiplicado por 2 !
	slice2 = append(slice2, 99)
	fmt.Printf("len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)

	/*****************************************************************
	   Trabalhando com Maps
	******************************************************************/

	// Outra maneira de declarar e já inicializar um map vazio !
	mapVazio := map[string]int{}
	fmt.Println(mapVazio)

	// Outra maneira de declarar e já inicializar um map com valores
	maps2 := map[string]int{"Paulo": 200, "Joca": 500, "Pedro": 7010}

	fmt.Println(maps2)

	// Exibindo apenas o salario da Paulo
	fmt.Println(maps2["Paulo"])

	// Removendo um elemento do mapa baseado na Chave
	delete(maps2, "Joca")
	fmt.Println(maps2)

	// Adicionando um elemento do mapa baseado na Chave
	maps2["Thiago"] = 900
	fmt.Println(maps2)

	percorrerMap(maps2)

}

func obterUltimoElementoArray(array [3]int) int {
	// Obtem o ultimo elemento do array
	return array[len(array)-1]
}

func percorrerArray(array [3]int) {
	fmt.Println("Percorrendo array ............")

	for i, v := range array {
		fmt.Printf("O valor do indice %d é %d", i, v)
		fmt.Println()
	}
}

func percorrerMap(salarios map[string]int) {
	for k, v := range salarios {
		fmt.Printf("O salario de %s é %d\n", k, v)
	}
}
