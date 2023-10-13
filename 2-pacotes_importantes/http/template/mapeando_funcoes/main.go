package main

import (
	/*
		Repare que aqui estamos usando html/template ao invés de text/template
		Quando trabalhamos com html, ele é melhor pis ja vem com uma série de validações que nos
		ajuda na prevenção de ataques.
	*/
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")

	/*
	 Na linha abaixo, estamos mapeando uma função do Go que converte string em maiscula
	 e através da chave ToUpper (poderia ser qualquer outro nome) vamos poder chamar ela no nosso arquivo html.
	*/
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
}
