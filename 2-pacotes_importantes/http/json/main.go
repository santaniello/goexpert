package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	jsonConta := serializeJson(conta)
	println(string(jsonConta))

	encodeJson(conta)

	jsonPuro := []byte(`{"numero":2,"saldo":200}`)
	contaDeserializada := deserializeJson(jsonPuro)
	fmt.Println(*contaDeserializada)

	jsonPuro2 := []byte(`{"Numero":3,"Saldo":300}`)
	contaDeserializada2 := deserializeJson(jsonPuro2)
	fmt.Println(*contaDeserializada2)

}

/*
json.Marshal:

Função: Converte uma estrutura de dados Go (por exemplo, struct, map, slice) em um slice de bytes representando o formato JSON.
Uso: É útil quando você quer obter uma representação JSON de um objeto para armazená-lo ou transmiti-lo, mas não necessariamente enviá-lo diretamente para um fluxo de saída.
*/
func serializeJson(conta Conta) []byte {
	json, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	return json
}

func deserializeJson(jsonConta []byte) *Conta {
	var conta Conta
	err := json.Unmarshal(jsonConta, &conta)
	if err != nil {
		println(err)
	}
	return &conta
}

/*
json.NewEncoder().Encode:

Função: Converte uma estrutura de dados Go diretamente para um formato JSON e o escreve em um io.Writer (por exemplo, um os.File, um http.ResponseWriter, etc.).
Uso: É útil quando você quer enviar uma representação JSON de um objeto diretamente para um fluxo de saída, como uma resposta HTTP ou um arquivo.

*/

func encodeJson(conta Conta) {
	err := json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}
}
