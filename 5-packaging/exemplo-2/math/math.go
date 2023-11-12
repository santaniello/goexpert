package math

var X string = "Hello World"

type Math struct {
	A int
	B int
	C int
}

type mathPrivate struct {
	a int
	b int
}

func NewMathPrivate(a, b int) mathPrivate {
	return mathPrivate{a: a, b: b}
}

func (m Math) Add() int {
	return m.A + m.B
}

func (m mathPrivate) Add() int {
	return m.a + m.b
}
