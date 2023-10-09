//No Go, a nomenclatura e organização de módulos segue um padrão relacionado ao controle de versão
//e ao caminho de importação.
//Aqui estão algumas diretrizes e práticas recomendadas:
//- URL do Repositório: O caminho de importação do módulo geralmente reflete a URL do repositório de controle de versão onde o módulo está hospedado. Por exemplo, se você tem um pacote no GitHub, o caminho do módulo será algo como:
//  github.com/seu_nome_de_usuario/nome_do_repositório
//- Versões Semânticas: Ao versionar módulos, Go usa versão semântica (semver). Se uma versão do seu módulo é v2 ou superior, a versão (v2, v3, etc.) é incluída como um sufixo no caminho do módulo. Por exemplo:
//  github.com/seu_nome_de_usuario/nome_do_repositório/v2
//Nome do Módulo: O nome do módulo em si (por exemplo, nome_do_repositório no exemplo acima) geralmente é escrito em "camelCase" se lacos_repeticao composto de várias palavras. No entanto, isso não é rigorosamente necessário; o importante é que o nome seja descritivo e siga as convenções do repositório de hospedagem (por exemplo, GitHub, GitLab, etc.) onde letras maiúsculas e minúsculas são tratadas de forma diferente.
//Nomes de Pacotes: Dentro de um módulo, você pode ter vários pacotes. A convenção para nomes de pacotes é usar letras minúsculas, sem espaços ou underscores. Se o nome do pacote consistir em várias palavras, elas são normalmente combinadas sem espaços ou separadas por um underscore (embora a combinação sem espaços seja mais comum).

module github.com/santaniello/fundamentos

go 1.20
