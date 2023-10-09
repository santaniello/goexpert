package main

import (
	"fmt"
	"strconv"
)

type Conta struct {
	Saldo int
}

func NewConta() *Conta {
	return &Conta{Saldo: 0}
}

// Repare que o valor do saldo fora dessa função não será afetado pois não estamos usando um ponteiro e sim uma copia da struct Conta
func (c Conta) Simular(valor int) int {
	c.Saldo += valor
	fmt.Printf("Valor do saldo dentro do método Simular %d \n", c.Saldo)
	return c.Saldo
}

func main() {
	fmt.Println("#################### Exemplo 1 ######################")

	a := 10
	imprimeValor(a)
	imprimeEnderecoDeMemoria(a)

	fmt.Println("***************************************************************************************")

	fmt.Println("alterando o valor do ponteiro  que aponta para a !")
	var ponteiro *int = &a
	*ponteiro = 20

	imprimeValorDePonteiro(ponteiro)
	imprimeValor(a)

	fmt.Println("***************************************************************************************")

	fmt.Println("alterando o valor do ponteiro chamado b que aponta para a !")
	b := &a
	*b = 30

	fmt.Printf("Valor de b %d \n", *b)
	imprimeValor(a)
	imprimeValorDePonteiro(ponteiro)

	fmt.Println("***************************************************************************************")

	fmt.Println("Alterando a copia da variavel a apenas dentro do método alteraOValorDaCopiaVariavel !")
	alteraOValorDaCopiaVariavel(a)
	imprimeValor(a)
	fmt.Printf("Valor de b %d \n", *b)
	imprimeValor(a)
	imprimeValorDePonteiro(ponteiro)

	fmt.Println("***************************************************************************************")

	fmt.Println("Alterando o valor da variavel a em todas as referências (ponteiros) !")
	alteraOValorDaVariavelEmTodasReferencias(&a)
	imprimeValor(a)
	fmt.Printf("Valor de b %d \n", *b)
	imprimeValor(a)
	imprimeValorDePonteiro(ponteiro)

	fmt.Println("#################### Exemplo 2 ######################")

	conta := NewConta()
	conta.Saldo = 100
	simulacao := conta.Simular(200)
	fmt.Printf("Valor da simulacao Bancaria %d \n", simulacao)

	// Repare que o valor do saldo não foi alterado !
	fmt.Printf("Valor do Saldo %d \n", conta.Saldo)

}

func imprimeValor(a int) {
	fmt.Println("Valor da variavel a = " + strconv.Itoa(a))
}

// & Comercial imprime o endereco
func imprimeEnderecoDeMemoria(a int) {
	fmt.Printf("Endereço da variável &a =  %p\n", &a)
}

// strconv.Itoa converte o valor de um ponteiro em uma string
func imprimeValorDePonteiro(ponteiro *int) {
	fmt.Println("Valor do *ponteiro para a = " + strconv.Itoa(*ponteiro))
}

func alteraOValorDoPonteiro(ponteiro *int) {
	*ponteiro = 50
	fmt.Println("Valor do ponteiro para a  após alteração = " + strconv.Itoa(*ponteiro))
}

func alteraOValorDaCopiaVariavel(a int) {
	a = 100
	fmt.Println("Valor da variavel local a após alteração = " + strconv.Itoa(a))
}

func alteraOValorDaVariavelEmTodasReferencias(a *int) {
	*a = 80
	fmt.Println("Valor da variavel global após alteração = " + strconv.Itoa(*a))
}
