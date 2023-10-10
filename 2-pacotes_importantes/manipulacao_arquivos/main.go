package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criação do Arquivo
	file := createFile()
	// O defer é um statement. ele atrasa a execução desse trecho de código sendo executado por ultimo.
	defer file.Close()

	// Gravando dados no arquivo
	tamanho := writeFile(file, "Escrevendo no arquivo")
	tamanho = writeFile(file, "Escrevendo no arquivo2")
	fmt.Printf("Arquivo criado com sucesso ! Tamanho: %d bytes\n", tamanho)

	readFile(file.Name())
	readFileChunks(file.Name())

}

func writeFile(file *os.File, message string) int {
	tamanho, err := file.Write([]byte(message + "\n"))
	if err != nil {
		panic(err)
	}
	return tamanho
}

func createFile() *os.File {
	file, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	return file
}

// Faz a leitura de um arquivo de uma unica vez !
func readFile(fileName string) {
	fmt.Println("Lendo arquivo .........")
	file, bytes := os.ReadFile(fileName)
	if bytes != nil {
		return
	}
	fmt.Println(string(file))
}

/*
A função abaixo lê um arquivo em partes (de 10 em 10 bytes)
*/
func readFileChunks(fileName string) {
	file, err := os.Open(fileName)
	// O defer é um statement. ele atrasa a execução desse trecho de código sendo executado por ultimo.
	defer file.Close()
	if err != nil {
		return
	}
	// Criamos um buffer para fazer a leitura do nosso arquivo
	reader := bufio.NewReader(file)
	// Criamos um slice de 10 bytes para armazenar os pedações que vão sendo lidos do nosso arquivo
	buffer := make([]byte, 10)
	for {
		position, err := reader.Read(buffer)
		if err != nil {
			return
		}
		// Exibe apenas a quantidade de bytes lido e que foram armazenadas no nosso slice
		fmt.Println(string(buffer[:position]))
	}
}
