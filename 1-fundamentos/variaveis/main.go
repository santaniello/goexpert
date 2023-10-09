/**
- Um package sempre terá o nome do diretório o qual está inserido com excessão do package main que é o
principal package do nosso sistema e é por onde nossa app vai comecar.

- Todos arquivos que estão dentro do mesmo diretório devem pertencer ao mesmo package.

- Tudo que está dentro do mesmo package (funções, váriaveis e etc...) pode ser acessado por todos os arquivos daquele diretório/package independentemente se são publicos ou privados.
*/

package main

import (
	"fmt"
)

func main() {
	/*****************************************************************
	   Imprimindo variaveis de outro arquivo do mesmo pacote
	******************************************************************/
	fmt.Println(a)
	// Inicializando uma variavel de outro arquivo do mesmo pacote
	b = true
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	// Imprimindo uma variavel que é do tipo Cep Criado por nós e que foi incializada em outro arquivo
	fmt.Println(g)

	/*****************************************************************
	  Maneira simplificada de declarar variaveis e tipos locais no Go
	  OBS: := Só funciona na primeira vez que a váriavel lacos_repeticao criada !
	******************************************************************/
	shortDeclaration := "Short"
	fmt.Println(shortDeclaration)
	shortDeclaration = "trocado valor do short"
	fmt.Println(shortDeclaration)

	// Aqui imprimoso tipo da variavel g que é o tipo main.Cep
	imprimeTipoVariavel(g)
}

func imprimeTipoVariavel(c Cep) {
	fmt.Printf("O tipo de g é %T", c)
	fmt.Println()
}
