package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	/* No Go não existe herança, porém, existe composição e nesse caso estamos usando umka composição anônima !
	   poderiamos dar um nome para ela exemplo: Address Endereco e na hora de chamar cliente.Address.Cidade ao invés de cliente.Endereco.Cidade
	*/
	Endereco
}

/*
*
O (c *Cliente) antes do nome da função Desativar() é um receiver em Go (Golang). Isso torna a função Desativar() um método associado ao tipo Cliente.
O * antes de Cliente indica que estamos lidando com um ponteiro para um Cliente, ou seja, uma referência a um objeto Cliente, e não uma cópia do objeto.

E se não tivessemos usando um ponteiro e sim uma copia do objeto cliente ?

Exemplo:

	func (c Cliente) Desativar() {
	    c.Ativo = false
	    fmt.Printf("O cliente %s foi desativado ", c.Nome)
	    fmt.Println()
	}

Sem o *, a função receberia uma cópia do valor Cliente, em vez de um ponteiro para o valor original. Isto significa que qualquer modificação feita em c dentro da função não afetaria o valor original do Cliente fora da função.

Por exemplo, na função acima, a linha c.Ativo = false alteraria o campo Ativo apenas da cópia local c e não do objeto Cliente original que chamou o método. Como resultado, o cliente original permaneceria ativo mesmo após a chamada do método.
*/
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado ", c.Nome)
	fmt.Println()
}
