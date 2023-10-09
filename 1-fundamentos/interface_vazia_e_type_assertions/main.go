package main

import "fmt"

func main() {
	/*
		Em Go, uma interface vazia, denotada como interface{}, não tem nenhum método especificado e qualquer tipo implementa a mesma.
		Aqui estão algumas razões para usar uma interface vazia:
		   - Armazenar Qualquer Tipo: Uma variável do tipo interface{} pode armazenar qualquer valor, independentemente do seu tipo. Isso torna a interface vazia útil para conter valores de tipos desconhecidos em tempo de compilação.
		   - Containers Genéricos: Estruturas de dados genéricas, como slices ou mapas, podem usar a interface vazia para armazenar elementos de qualquer tipo.
		   - Decode/Unmarshal de Dados: Ao lidar com dados desconhecidos ou variáveis, como ao decodificar JSON, a interface vazia é frequentemente usada para receber os dados.
		   - Type Assertion e Type Switches: Embora você possa armazenar qualquer valor em uma interface vazia, frequentemente você precisará saber o tipo real do valor e talvez convertê-lo de volta para esse tipo. Isso é feito usando type assertions e type switches.
		   - Comunicação entre partes desacopladas do código: Em sistemas complexos, onde diferentes partes do código não devem ter dependências rígidas entre si, a interface vazia pode ser usada para enviar/receber dados de maneira mais genérica.
		Apesar de sua utilidade, o uso excessivo da interface vazia pode levar a problemas. O principal é a perda de informações de tipo em tempo de compilação, o que pode levar a erros em tempo de execução que poderiam ter sido evitados. Sempre que possível, é melhor usar tipos concretos ou interfaces definidas. A interface vazia deve ser usada com cautela e apenas quando o benefício superar os potenciais problemas.

		O uso de interfaces vazias é frequente em sistemas legados e vem sendo substituido gradativamento pelo uso de Generics
	*/
	var x interface{} = 10
	var y interface{} = "Hello"
	showType(x)
	showType(y)
	usarTypeAssertions(x)
	usarTypeAssertions(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)

}

func usarTypeAssertions(t interface{}) {
	/*
		Na linha de baixo fazemos um type assertation para saber o tipo real do valor que a interface vazia esta armazenando.
		A variavel resString contém o valor após a conversão para string e okString tem um boolean indicando true ou false para caso a conversão tenha sido realizada ou não.
	*/
	resString, okString := t.(string)
	fmt.Printf("O valor de resString é %v e o resultado de okString é %v\n", resString, okString)

	/*
	 Na linha de baixo fazemos um type assertation para saber o tipo real do valor que a interface vazia esta armazenando.
	 A variavel resInt contém o valor após a conversão para int e okInt tem um boolean indicando true ou false para caso a conversão tenha sido realizada ou não.
	*/
	resInt, okInt := t.(int)
	fmt.Printf("O valor de resInt é %v e o resultado de okInt é %v\n", resInt, okInt)

}
