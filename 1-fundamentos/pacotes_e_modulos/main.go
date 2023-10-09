package main

import (
	"fmt"
	"github.com/santaniello/fundamentos/pacotes_e_modulos/matematica"
)

func main() {

	/*
	 funções, variaveis, structs que começam com letra maiuscula são acessiveis fora do pacote as quais pertencem pois seria como se
	 o modficiador de acesso delas fosse publico
	*/
	somaPublico := matematica.SomaPublico(10, 20)
	fmt.Printf("Resultado da Soma Publico: %v\n", somaPublico)

	/*
		funções, variaveis, structs que começam com letra minuscula são acessiveis APENAS dentro do pacote as quais elas estão inseridas.
		Se tentarmos chamar uma função privada por exemplo do pacote matematica, receberemos o seguinte erro:
		./main.go:12:28: undefined: matematica.somaprivado
	*/

	//somaPrivado := matematica.somaprivado(10, 20)

	fmt.Printf("Variavel Publica : %v\n", matematica.VariavelPublica)

	/*
	  Não conseguimos usar a variavel com o escopo privado aqui.
	*/
	//fmt.Printf("Variavel Publica : %v\n", matematica.variavelPrivada)

}
