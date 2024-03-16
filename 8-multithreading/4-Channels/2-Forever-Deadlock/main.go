package main

import "fmt"

func main() {
	//SimulandoDeadlockComForever()
	//SimulandoDeadlockComForever2()
	//SimulandoDeadlockComForever3()
	ResolvendoDeadlock()
}

func SimulandoDeadlockComForever() {
	forever := make(chan bool) // cria um canal vazio que nunca será fechado

	<-forever // bloqueia a thread principal até que o canal seja fechado, porém como o canal nunca será fechado, a thread principal  nunca será desbloqueada gerando um deadlock
}

func SimulandoDeadlockComForever2() {
	forever := make(chan bool) // cria um canal vazio que nunca será fechado

	go func() { // aqui criamos uma goroutine que vai executar um loop até 10, porém em nenhum momento ela fecha o canal forever, gerando um deadlock
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	<-forever // bloqueia a thread principal até que o canal seja fechado, porém como o canal nunca será fechado, a thread principal  nunca será desbloqueada gerando um deadlock
}

func SimulandoDeadlockComForever3() {
	forever := make(chan bool) // cria um canal vazio que nunca será fechado

	forever <- true // o canal forever gera um deadlock nesse caso pois deveria ser preenchido em outra goroutine e não na mesma onde será lido

	<-forever
}

func ResolvendoDeadlock() {
	forever := make(chan bool) // cria um canal vazio que nunca será fechado

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		forever <- true // aqui enviamos uma mensagem para o canal forever, desbloqueando a thread principal
	}()

	<-forever // bloqueia a thread principal até que o canal seja fechado, porém como o canal nunca será fechado, a thread principal  nunca será desbloqueada gerando um deadlock
}
