## GO project layout

Segue um repo que é um modelo interessante de estrutura que podemos seguir para nossoas projetos em Golang:

https://github.com/golang-standards/project-layout

## Pastas importantes

- internal: Essa pasta contém os arquivos que são exclusivos do nosso negócio e não queremos compartilhar através do go mod.
- pkg: Ao contrário da pasta internal, a pasta pkg contém libraries que nós podemos compartilhar com outros projetos.
- cmd/nome_da_aplicacao_folder: Diretório onde vai ficar o seu executável. Aqui fica localizado o arquivo main.go
- configs: contém os arquivos para o bootstrap da aplicação
- test: Aplicativos de teste externos adicionais e dados de teste.
- api: Guarda as especificações da nossa api (swagger).
