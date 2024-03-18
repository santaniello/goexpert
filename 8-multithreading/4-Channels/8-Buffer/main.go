package main

/*
Devemos utilizar o buffer quando queremos que o canal possa armazenar mais de um valor sem bloquear a execução do programa, porém, para que tenhamos
um aumento de performance, devemos utilizar o buffer com cautela, pois o mesmo pode aumentar o consumo de memória do programa.
O ideal é que o buffer seja utilizado em casos onde temos um grande número de produtores e consumidores, ou quando o tempo de processamento  de um
dos lados do canal é muito maior que o outro.

Faça sempre um benchmark e ue com parsimônia.
*/

func main() {
	ch := make(chan string, 2) // Aqui exemplificamos a criação de um canal com buffer de tamanho 2...

	ch <- "hello" // podemos adicionar 2 valores no canal sem bloquear a execução do programa
	ch <- "world"
	println(<-ch) // Aqui estamos lendo o primeiro valor do canal
	println(<-ch) // Aqui estamos lendo o segundo valor do canal

}
