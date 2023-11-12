package main

import (
	"github.com/google/uuid"
	"github.com/santaniello/packaging/exemplo-2/math"
)

func main() {
	m := math.NewMathPrivate(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
