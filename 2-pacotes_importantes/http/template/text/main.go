package main

import (
	"os"
	"text/template"
)

type Pessoa struct {
	Nome  string
	Idade int
}

/*
Em Go (também conhecido como Golang), "templates" são uma forma de representar dados em um formato específico usando uma string ou arquivo de modelo. O pacote padrão text/template e seu análogo html/template fornecem uma maneira de gerar strings personalizadas com base em dados de entrada.

Os templates no Go são úteis em várias situações, como:

Renderizar HTML: Para aplicações web, você pode usar o pacote html/template para garantir que os dados sejam exibidos de forma segura, evitando ataques como Cross-site Scripting (XSS).

Configurações Dinâmicas: Para criar arquivos de configuração a partir de uma estrutura de dados.

Saída Formatada: Como gerar relatórios ou exibir dados no terminal de forma personalizada.

Neste exemplo, o template Olá {{.Nome}}, você tem {{.Idade}} anos! é usado para formatar uma mensagem com base nos campos Nome e Idade da struct Pessoa.

As ações (como {{.Nome}}) no template são substituídas pelos valores correspondentes do dado fornecido (neste caso, a instância p da struct Pessoa).

*/

func main() {
	executarTemplate()
	executarTemplateComMust()
}

func executarTemplate() {
	t := template.New("nomeEIdade")
	t, _ = t.Parse("Olá {{.Nome}}, você tem {{.Idade}} anos!\n")
	p := Pessoa{Nome: "Alice", Idade: 30}
	t.Execute(os.Stdout, p)
}

/*
 A função abaixo faz exatamente a mesma coisa que a função acima, porém,
 A função template.Must é uma função auxiliar disponível tanto no pacote text/template quanto no html/template em Go.
 Ela é usada para encapsular a chamada para uma função que retorna um valor e um erro, e se um erro for retornado, Must irá causar um panic.

*/

func executarTemplateComMust() {
	temp := template.Must(template.New("nomeEIdade").Parse("Olá {{.Nome}}, você tem {{.Idade}} anos!\n"))
	p := Pessoa{Nome: "Alice", Idade: 30}
	err := temp.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
