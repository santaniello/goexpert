package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	SimularApiComProblemaDeConcorrencia()
	//CorrigindoApiComProblemaDeConcorrenciaUsandoMutex()
	//CorrigindoApiComProblemaDeConcorrenciaUsandoVariaveisAtomicas()

}

func SimularApiComProblemaDeConcorrencia() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

func CorrigindoApiComProblemaDeConcorrenciaUsandoMutex() {
	mut := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mut.Lock() // Não deixa ninguém acessar a variável number enquanto estiver sendo usada
		number++
		mut.Unlock() // Libera a variável number para ser usada por outra goroutine
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

func CorrigindoApiComProblemaDeConcorrenciaUsandoVariaveisAtomicas() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1) // Adiciona 1 na variável number atomicamente
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número: %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
