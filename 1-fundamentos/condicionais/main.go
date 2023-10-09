package main

func main() {
	a := 1
	b := 2

	if a > b {
		println(a)
	} else { // No Go, nós não temos else if
		println(b)
	}

	// && é o and
	if a > b && b > 0 {
		println(a)
	} else { // No Go, nós não temos else if
		println(b)
	}

	// || é o or
	if a > b || b > 0 {
		println(a)
	} else { // No Go, nós não temos else if
		println(b)
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("Default")
	}
}
