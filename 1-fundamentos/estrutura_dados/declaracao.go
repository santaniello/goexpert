package main

/*
*
Array em Go:
  - Uma coleção fixa de elementos do mesmo tipo.
  - O tamanho do array é parte de sua definição, então você não pode alterar seu tamanho após sua criação.
*/
var meuArray [3]int

/*
Slice em Go:
  - Seria uma espécie de array dinâmico que pode ter seu tamanho redimensionado.
  - Slices não possuem tamanho definido
  - A maioria das operações em Go que você fará com coleções de elementos será com slices e não arrays.

O que é a função make usada para inicializar o slice ?

A função make é uma função embutida (built-in) no Go que é usada para alocar e inicializar slices, mapas e canais, que são três dos tipos de dados
mais importantes do Go que têm uma natureza dinâmica e precisam de alocação especial.

A expressão slice = make([]int, 0) faz o seguinte:

A chamada à função embutida make cria uma slice do tipo int com tamanho inicial de 0,
o que significa que ela não possui elementos imediatamente após sua criação.
Como não especificamos uma capacidade diferente, ela será igual ao seu tamanho, que é 0.
Assim, ao adicionar elementos subsequentemente (por exemplo, usando append), a slice será realocada.

OBS: Se lacos_repeticao trabalhar com uma quantidade de dados muito grande, tente inicializar o slice com um tamanho
proóximo ao que vc vai trabalhar para evitar que ele fique se auto redimensionando e isso afete
a performance da sua aplicação
*/
var slice []int = make([]int, 0)

/*
*
Map em Go:
  - Uma coleção de pares chave-valor.
  - As chaves são únicas, e cada chave mapeia para exatamente um valor.
  - É similar ao conceito de dicionário ou tabela hash em outras linguagens.
  - Ele não é ordenado !

OBS: A função make é uma função embutida (built-in) no Go que é usada para alocar e inicializar slices, mapas e canais, que são três dos tipos de dados
mais importantes do Go que têm uma natureza dinâmica e precisam de alocação especial.
*/
var maps map[string]int = make(map[string]int)
