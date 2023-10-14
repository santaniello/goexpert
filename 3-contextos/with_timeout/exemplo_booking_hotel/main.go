package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)

}

/*
O select é uma estrutura de controle em Go (também conhecido como Golang) usada principalmente para lidar com múltiplos canais. Ela permite que o programa espere até que uma das comunicações possa seguir em frente, o que é útil para lidar com operações concorrentes.

Na função bookHotel(ctx context.Context) que você forneceu, o select é usado para esperar por um dos seguintes cenários:

O ctx.Done() é chamado: Isso geralmente ocorre quando o context.Context que é passado para a função é cancelado ou excede seu tempo limite. Quando isso acontece, a mensagem "Hotel booking cancelled. Timeout reached." é impressa, indicando que o processo de reserva foi cancelado devido ao tempo limite.

A função time.After(1 * time.Second) envia um sinal: Isso ocorre após 1 segundo. Se isso acontecer antes de ctx.Done() ser chamado, a mensagem "Hotel booked." é impressa, indicando que a reserva foi feita.

Neste exemplo, o select serve para lidar com um cenário onde você deseja tentar reservar um hotel, mas apenas se a operação puder ser concluída em um certo período de tempo (neste caso, antes que o contexto seja finalizado ou em até 1 segundo). Se qualquer uma das condições for atendida primeiro, a operação correspondente será executada.

Em resumo, o select em Go é uma ferramenta poderosa para trabalhar com concorrência, permitindo que você espere e responda a várias possíveis comunicações de canais de forma eficiente.
*/
func bookHotel(ctx context.Context) {
	select {
	/*
		Em Go (Golang), o símbolo <- é utilizado principalmente no contexto de operações com canais (channels). Canais são uma ferramenta poderosa em Go para comunicação entre goroutines, permitindo a troca segura de dados entre elas.
		O símbolo <- é usado tanto para enviar quanto para receber valores de um canal.
	*/
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
