package main

import (
	"fmt"
	"github.com/santaniello/packaging/exemplo-1/math"
)

func main() {
	// Usando Math Publico
	m := math.Math{A: 1, B: 2}
	m.C = 3
	fmt.Println(m.Add())

	// Usando Math Privado
	mPrivate := math.NewMathPrivate(1, 2)
	/**
	Repare que só temos acesso ao método Add que exportamos !
	Se quisessemos ter acesso as variaveis, precisariamos criar métodos publicos para acessa-las (getters e setters)
	*/
	fmt.Println(mPrivate.Add())
	// Aqui conseguimos acessar uma variavel publica
	fmt.Println(math.X)
}
