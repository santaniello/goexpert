package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		// Repare que aqui estamos passando o nosso ResponseWriter como saida e não o nosso console como em outros exemplos...
		err := t.Execute(rw, Cursos{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", mux)
}
