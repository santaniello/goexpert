package main

import "fmt"

func SomaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

/*
Repare que no meu método SomaInterfaceVazia, eu posso receber um map com quaisquer valores (int, float, string) e realizar a soma.
Antes do Generic, era assim que as coisas era feitas.
*/
func SomaInterfaceVazia(m map[string]interface{}) float64 {
	var soma float64
	for _, v := range m {
		switch t := v.(type) {
		case int:
			soma += float64(t)
		case float64:
			soma += t
		default:
		}
	}
	return soma
}

/*******************
 	GENERICS
********************/

/*
Abaixo criamos uma Generic chamada T que aceita tanto o tipo int quanto o float e pode devolver um T que pode ser tanto int quanto float
*/
func SomaComGenerics[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

/*
Abaixo nós criamos aquilo que chamamos de constraint (restrição) para usar Generics.
O tipo Number é uma interface que pode ser satisfeita por valores que são do tipo int ou float64.
A restrição de conjunto de tipos permite definir de forma mais clara e concisa quais tipos exatos são permitidos, em vez de depender apenas de métodos de interface para restringir os tipos.
OBS: O Go ja tem uma série de constraint pré definidas que você pode usar ao invés de criar as suas https://pkg.go.dev/golang.org/x/exp/constraints#pkg-overview

O sinal ~ permite que você trabalhe com os tipos subjacentes quando você está lidando com restrições de conjunto de tipos em generics em Go,
e é por isso que você pode usar MyNumber em uma função que espera um int pois MyNumber é um tipo definido pelo usuário (ou "type alias"),
e seu tipo subjacente é int.
*/

type Number interface {
	~int | ~float64
}

func SomaComGenericsUsandoConstraint[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

/*
Método que compara dois generic types usando a constraint comparable
https://aprendagolang.com.br/2022/04/06/o-que-e-e-como-usar-a-nova-constraint-comparable/
*/
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	mInt := map[string]int{"a": 100, "b": 200}
	println(SomaInt(mInt))
	mFloat := map[string]float64{"a": 100.0, "b": 200.0}
	println(SomaFloat(mFloat))
	mInterfaceVazia := map[string]interface{}{"a": 100, "b": 200.0, "c": "Teste"}
	println(SomaInterfaceVazia(mInterfaceVazia))

	fmt.Println("Usando Generics ...............")
	println(SomaComGenerics(mInt))
	println(SomaComGenerics(mFloat))

	println(SomaComGenericsUsandoConstraint(mInt))
	println(SomaComGenericsUsandoConstraint(mFloat))

	mMyNumber := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	println(SomaComGenericsUsandoConstraint(mMyNumber))

	println(Compara(10, 10))
}
