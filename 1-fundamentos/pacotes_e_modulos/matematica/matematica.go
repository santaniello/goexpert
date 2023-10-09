package matematica

var VariavelPublica = 10
var variavelPrivada = 5

func SomaPublico[T int | float64](a, b T) T {
	return a + b
}

func somaprivado[T int | float64](a, b T) T {
	return a + b
}
