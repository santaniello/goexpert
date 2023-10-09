package main

/*
Qualquer Struct que tiver o método Desativar(), estará implementando a interface Pessoa
No nosso caso, quem esta implementando é a struct Cliente !
Quem faz essa inferência é o compilador do Go
*
*/
type Pessoa interface {
	Desativar()
}
