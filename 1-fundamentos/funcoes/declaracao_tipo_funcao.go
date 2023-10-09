package main

import "fmt"

// Define o tipo de função para filtragem: aceita um inteiro e retorna um booleano.
type filtroFunc func(int) bool

// Função que verifica se um número é par.
func ehPar(num int) bool {
	return num%2 == 0
}

// Função que verifica se um número é ímpar.
func ehImpar(num int) bool {
	return num%2 != 0
}

/*
*
Função que filtra uma slice de inteiros com base em um critério de filtragem fornecido.
Repare que aqui recebemos uma função como parâmetro

Como o Go sabe que a função ehPar e a função ehImpar são do tipo filtroFunc ?

Em Go, a determinação de tipos para funções é feita com base na assinatura da função, não necessariamente no nome do tipo definido.

Quando definimos um tipo de função, estamos definindo uma assinatura específica para essa função. Por exemplo, ao definir:

type filtroFunc func(int) bool

Estamos dizendo que filtroFunc é um tipo de função que aceita um único argumento do tipo int e retorna um bool.

Ambas as funções ehPar e ehImpar têm essa assinatura.

O compilador Go, ao verificar o código, reconhece que as funções ehPar e ehImpar correspondem à assinatura definida por filtroFunc. Portanto, elas podem ser usadas onde quer que uma filtroFunc seja esperada, sem a necessidade de uma conversão explícita.

É a assinatura da função (tipos de argumentos e tipo de retorno) que determina a compatibilidade, não necessariamente o nome do tipo. Se tivéssemos outras funções com a mesma assinatura (um único argumento int e um retorno bool), elas também seriam consideradas do tipo filtroFunc, mesmo que nunca as tenhamos explicitamente associado a esse tipo.
*/
func filtrar(nums []int, f filtroFunc) []int {
	var resultado []int
	for _, v := range nums {
		if f(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

func imprimeNumerosParesEhImpares() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Filtra números pares.
	pares := filtrar(numeros, ehPar)
	fmt.Println(pares) // [2 4 6 8 10]

	// Filtra números ímpares.
	impares := filtrar(numeros, ehImpar)
	fmt.Println(impares) // [1 3 5 7 9]
}
