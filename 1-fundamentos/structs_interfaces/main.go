package main

import "fmt"

func main() {
	// Instanciando uma struct
	felipe := Cliente{
		Nome:  "Felipe",
		Idade: 30,
		Ativo: true,
	}

	felipe.Ativo = true
	fmt.Println(felipe.Nome)
	fmt.Println(felipe.Ativo)
	felipe.Endereco.Cidade = "São Paulo"
	felipe.Endereco.Estado = "São Paulo"
	felipe.Endereco.Numero = 10
	felipe.Endereco.Logradouro = "Teste"
	Desativacao(&felipe)
	fmt.Println(felipe)
}

/*
Podemos passar todas as implementações de Pessoa (No caso só temos o cliente atualmente)

O parâmetro pessoa espera algo que satisfaça a interface Pessoa, não importa se é um tipo de valor ou um ponteiro. Na sua situação, é um ponteiro para Cliente (*Cliente) que satisfaz essa interface, então é isso que você passa para a função.

Você não precisa (e nem pode) usar *Pessoa na assinatura da função porque Pessoa é uma interface, e você não cria ponteiros para interfaces da mesma maneira que faz com tipos concretos. Em vez disso, o próprio tipo concreto (ou seu ponteiro) que você passa para a função determina como ele satisfaz a interface.
*
*/
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}
