package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	// Criamos um slice com o nome dos arquivos que queremos compor ...
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	/*
	 Aqui vamos utilizar o arquivo content como base e parsear todos os nossos arquivos html...
	 A função ParseFiles é uam função variádica, ou seja, recebe inumeros argumentos do tipo string e a sintaxe templates... faz com que passemos cada item do nosso slice como um parâmetro para essa função
	*/
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
}
