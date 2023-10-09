package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(1, 2))
	// Obtendo o retorno de uma função com multiplos retornos
	valor, err := sum3(1, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}

	// Usando a função variadica
	fmt.Println(sum4(1, 2, 5, 6, 7, 8, 9, 10))

	/*
	  Abaixo criamos uma função anônima no GO

	*/
	totalFuncaoAnonima := func(multiplicador int) int {
		return sum2(2, 2) * multiplicador
	}(2)

	fmt.Println(totalFuncaoAnonima)

	/*
	 Abaixo nós obtivemos o retorno da funcao contador que é outra função anonima
	 e depois executamos esse retorno
	*/
	funcaoRetorno := contador()
	fmt.Println(funcaoRetorno())

	/*
	 Abaixo testamos o uso de uma função que foi criada com um tipo epassada como parâmetro para outra função
	*/
	imprimeNumerosParesEhImpares()
}

func sum(a int, b int) int {
	return a + b
}

/*
Se os dois parâmetros tiverem o mesmo tipo, podemos ocultar o tipo do primeiro parâmetro
*/
func sum2(a, b int) int {
	return a + b
}

/*
Uma função pode retornar mais de um valor !
No caso abaixo, ela retornará os tipos (int, error)
*/
func sum3(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("erro, A soma é maior que 50")
	}
	return a + b, nil
}

/*
Aqui temos um exemplo de função variadica a qual podemos passar uma infinidade parâmetros
do tipo int !
*/
func sum4(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

/*
Em Go, as funções são cidadãos de primeira classe, o que significa que podem ser
atribuídas a variáveis, passadas como argumentos e retornadas de outras funções.
Isso, combinado com a capacidade de uma função interna capturar e usar variáveis
do seu ambiente circundante, permite a criação de closures.
*/
func contador() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
