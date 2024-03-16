package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID      int64
	Content string
}

// O objetivo desse exemplo é mostrar como podemos usar o select para receber mensagens de diferentes canais
func main() {
	kafkaChannel := make(chan Message)
	rabbitMQChannel := make(chan Message)
	var id int64 = 0
	go ProduceKafkaMessage(kafkaChannel, id)
	go ProduceRabbitMQMessage(rabbitMQChannel, id)
	recebendoMensagemDeDiferentesCanais(kafkaChannel, rabbitMQChannel)
}

func recebendoMensagemDeDiferentesCanais(kafkaChannel, rabbitMQChannel chan Message) {
	for {
		// O select é uma estrutura de controle que permite que um programa aguarde por várias operações de comunicação.
		select {
		// Aqui estamos recebendo mensagens de 2 canais diferentes
		case msg := <-kafkaChannel:
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg.ID, msg.Content)
		case msg := <-rabbitMQChannel:
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.ID, msg.Content)
		// O timeout é uma forma de evitar que o select fique bloqueado indefinidamente, caso nenhum dos canais esteja pronto para enviar uma mensagem após 5 segundos, o select irá executar o case do timeout
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			// O default é executado caso nenhum dos canais esteja pronto para enviar uma mensagem
			//default:
			//	fmt.Println("Default")
		}
	}
}

// ProduceKafkaMessage Simulando a produção de mensagens para o Kafka
func ProduceKafkaMessage(kafkaChannel chan<- Message, id int64) {
	for {
		time.Sleep(2 * time.Second)
		kafkaChannel <- Message{ID: atomic.AddInt64(&id, 1), Content: "Hello World from Kafka"}
	}
}

// ProduceRabbitMQMessage Simulando a produção de mensagens para o RabbitMQ
func ProduceRabbitMQMessage(RabitMQChannel chan<- Message, id int64) {
	for {
		time.Sleep(3 * time.Second)
		RabitMQChannel <- Message{ID: atomic.AddInt64(&id, 1), Content: "Hello World from RabbitMQ"}
	}
}
