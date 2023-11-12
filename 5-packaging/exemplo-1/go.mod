// O nome do módulo normalmente é composto por: repositorio onde o modulo ficará hospedado / owner do repo / nome do projeto
// Essa url será usada pelo Go pois sempre que referenciarmos ela em um projeto, o Go vai tentar baixar essa dependência usando o nome do módulo como URL
// o nome do modulo poderia ser: meumodulo/packaging porém, isso não é indicado pelo motivo listado acima (ele tentará fazer o download usando o nome do modulo como base e não irá conseguir).
module github.com/santaniello/packaging/exemplo-1

go 1.21
